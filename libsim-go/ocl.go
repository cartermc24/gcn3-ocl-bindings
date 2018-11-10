package main

/*
 */
import "C"
import (
	"container/list"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"reflect"
	"unsafe"
	"os"
	"os/user"
	"os/exec"
	"bytes"
	"encoding/hex"

	"gitlab.com/akita/gcn3/driver"
	"gitlab.com/akita/gcn3/insts"
	"gitlab.com/akita/gcn3/kernels"
	"gitlab.com/akita/gcn3/platform"
)

// Globals
var sim_driver *driver.Driver
var program_map map[int]*CLProgram
var buffer_map map[int]*driver.GPUPtr // TODO update to just uint32
var kernel_map map[int]*CLKernel
var program_idx int = 0
var buffer_idx int = 1 // Start at one to avoid conflict with C's NULL
var kernel_idx int = 0

// Hardcoded value
var context_id int = 8
var command_queue_id int = 12
var platform_id int = 0
var device_id int = 0

//
// Data Model
//
type CLProgram struct {
	program_string string
	program        []byte
}

type CLKernel struct {
	arg_list *list.List
	kernel   *insts.HsaCo
}

/*
   CLKernelArg Types
   0 -> GlobalPtr/Primitive
   1 -> LocalPtr
*/
type CLKernelArg struct {
	idx      int
	size     int
	ptr_val  uint64
	arg_type uint8
}

/*
func createKernelArgInterface(args []CLKernelArg) interface{} {
    f := make([]reflect.StructField, len(args))
    for i, _ := range args {
        arg_type := args[i].arg_type

        if arg_type == 0 {
            f[i].Type = reflect.TypeOf((*driver.GPUPtr) (nil)).Elem()
        } else {
            f[i].Type = reflect.TypeOf((*driver.LocalPtr) (nil)).Elem()
        }

        //f[i].Type = reflect.TypeOf(u)
        f[i].Anonymous = true
    }

    r := reflect.New(reflect.StructOf(f)).Elem()
    for i, u := range args {
        r.Field(i).Set(reflect.ValueOf(u))
    }
    return r.Addr().Interface()
*/

//
// Helpers
//
//export initializeSimulator
func initializeSimulator() {
	buffer_map = make(map[int]*driver.GPUPtr)
	kernel_map = make(map[int]*CLKernel)
	program_map = make(map[int]*CLProgram)
	_, _, sim_driver, _ = platform.BuildEmuPlatform()
	fmt.Println("[ocl-wrapper] Simulator Initialized\n")
}

// Convert CLKernelArg struct slice into raw bytes
// so that the simulator can handle it
func convertArgsToBytes(cl_kernel_args []CLKernelArg) []byte {
	var all_args []byte
	for i, kernel_arg := range cl_kernel_args {
		arg_bytes := make([]byte, unsafe.Sizeof(kernel_arg.ptr_val))
		binary.LittleEndian.PutUint64(arg_bytes, uint64(kernel_arg.ptr_val)) // FIXME assumes 64-bit platform	
		fmt.Printf("[ocl-wrapper] Argument %i has bytes: [%s]\n", i, hex.Dump(arg_bytes))
		all_args = append(all_args, arg_bytes...)
	}

	fmt.Printf("[ocl-wrapper] Arguments passed to kernel [%s]", hex.Dump(all_args))

	return all_args
}

//
// OpenCL-specific version of sim driver functions
//
func driverUpdateLDSPointers(
	d *driver.Driver,
	co *insts.HsaCo,
	clKernelArgs []CLKernelArg,
) {
	ldsSize := uint32(0)
	for _, kernelArg := range clKernelArgs {
		// Local Ptr
		if kernelArg.arg_type == 1 {
			kernelArg.ptr_val = uint64(ldsSize)
			ldsSize += uint32(kernelArg.ptr_val)
		}
	}
	co.WGGroupSegmentByteSize = ldsSize
}

//
// OpenCL API
//
//export gcn3GetPlatformIDs
func gcn3GetPlatformIDs() int {
	return platform_id
}

//export gcn3GetDeviceIDs
func gcn3GetDeviceIDs() int {
	return device_id
}

//export gcn3CreateContext
func gcn3CreateContext() int {
	return context_id
}

//func gcn3CreateCommandQueue

// Returns Kernel ID
//export gcn3CreateProgramWithSource
func gcn3CreateProgramWithSource(context int, program_string string) int {
	if context != context_id {
		fmt.Println("[ocl-wrapper] Error creating program from source: invalid context")
		return -34 // CL_INVALID_CONTEXT
	}

	cl_program := CLProgram{}
	cl_program.program_string = program_string

	program_map[program_idx] = &cl_program

	program_idx += 1

	fmt.Printf("[ocl-wrapper] Created program with ID %v\n", program_idx-1)

	return program_idx - 1
}

//export gcn3BuildProgram
func gcn3BuildProgram(program_id int) int {
	// Write program source to file
	fmt.Printf("[ocl-wrapper] Writing CL source to temporary file\n")
	program_bytes := []byte(program_map[program_id].program_string)
	write_err := ioutil.WriteFile("/tmp/prog.cl", program_bytes, 0644)
	if write_err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: could not write CL source to /tmp\n")
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

	// Run compiler
	fmt.Printf("[ocl-wrapper] Running clang-ocl\n")
	usr, usr_err := user.Current()
	if usr_err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: unable to get user information to locate compiler\n")
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}
	compiler_root := usr.HomeDir + "/.clangocl/clang-ocl"

	compiler := exec.Command(compiler_root, "-mcpu=gfx803", "-o", "/tmp/prog.hsaco", "/tmp/prog.cl")

	var stderr bytes.Buffer
	compiler.Stderr = &stderr
	err := compiler.Run()

	if err != nil && err.Error() == "exit status 1" {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: clang-ocl reported compiler errors:\n")
		fmt.Fprintf(os.Stderr, stderr.String())
		fmt.Fprintf(os.Stderr, "\n")
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: could not invoke compiler, ensure clang-ocl exists at ~/.clangocl/clang-ocl, error: %s\n", err)
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

	fmt.Printf("[ocl-wrapper] CL source successfully compiled\n")
	hsacoBytes, err := ioutil.ReadFile("/tmp/prog.hsaco")
	if err != nil {
		fmt.Printf("[ocl-wrapper] Error building program: %v\n", err)
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

	program_map[program_id].program = hsacoBytes

	fmt.Printf("[ocl-wrapper] Built program with ID %v\n", program_id)
	return 1 // CL_SUCCESS
}

//export gcn3CreateKernel
func gcn3CreateKernel(program_id int, kernel_name string) int {

	fmt.Println("[ocl-wrapper] Attempting to create kernel")
	program := program_map[program_id].program

	cl_kernel := CLKernel{}
	kernel := kernels.LoadProgramFromMemory(program, kernel_name)

	if kernel == nil {
		fmt.Println("[ocl-wrapper] Unable to get kernel from program")
		return -46 // CL_INVALID_KERNEL_NAME
	}

	cl_kernel.kernel = kernel
	cl_kernel.arg_list = list.New()
	kernel_map[kernel_idx] = &cl_kernel

	kernel_idx += 1

	fmt.Printf("[ocl-wrapper] Created kernel with name: %v, ID: %v\n", kernel_name, kernel_idx-1)
	return kernel_idx - 1
}

// Returns Buffer ID
// Size in bytes
//export gcn3CreateBuffer
func gcn3CreateBuffer(context int, size int) int {
	if context != context_id {
		return -34 // CL_INVALID_CONTEXT
	}

	new_buffer := sim_driver.AllocateMemory(uint64(size))
	buffer_map[buffer_idx] = &new_buffer

	buffer_idx += 1

	fmt.Printf("[ocl-wrapper] Allocated buffer of size: %v, wrapper ID: %v, sim addr: %v\n", size, buffer_idx-1, new_buffer)
	return buffer_idx - 1
}

//export gcn3EnqueueWriteBuffer
func gcn3EnqueueWriteBuffer(buffer int, size int, ptr unsafe.Pointer) int {
	sim_buffer := buffer_map[buffer]

	ptr_bytes := C.GoBytes(ptr, C.int(size))

	sim_driver.MemoryCopyHostToDevice(*sim_buffer, ptr_bytes)

	//fmt.Printf("[ocl-wrapper] Wrote data to device: [%s] @ region: %02x\n", hex.Dump(ptr_bytes), *sim_buffer)

	fmt.Printf("[ocl-wrapper] Enqueued Write Buffer for buffer ID %v\n", buffer)
	return 0 // CL_SUCCESS
}

//export gcn3EnqueueReadBuffer
func gcn3EnqueueReadBuffer(buffer int, size int, ptr unsafe.Pointer) int {
	sim_buffer := buffer_map[buffer]

	ptr_bytes := C.GoBytes(ptr, C.int(size))

	//fmt.Printf("[ocl-wrapper] Fetched data from device: [%s] @ region %02x\n", hex.Dump(ptr_bytes), *sim_buffer)

	sim_driver.MemoryCopyDeviceToHost(ptr_bytes, *sim_buffer)

	fmt.Printf("[ocl-wrapper] Enqueued Read Buffer for buffer ID %v\n", buffer)

	return 0 // CL_SUCCESS
}

//export gcn3SetKernelArg
func gcn3SetKernelArg(kernel int, arg_idx int, size int, ptr unsafe.Pointer) int {
	cl_kernel := kernel_map[kernel]
	arg_list := cl_kernel.arg_list
	cl_kernel_arg := CLKernelArg{}
	cl_kernel_arg.idx = arg_idx
	cl_kernel_arg.size = size
	//cl_kernel_arg.bytes = *((*[size]byte) (ptr))
	//cl_kernel_arg.bytes = C.GoBytes(ptr, C.int(size))
	cl_kernel_arg.ptr_val = *(*uint64)(ptr) //uint64(uintptr(ptr))

	test := unsafe.Sizeof(reflect.TypeOf((int)(0)))

	if uintptr(size) > test {
		cl_kernel_arg.arg_type = uint8(1)
	} else {
		cl_kernel_arg.arg_type = uint8(0)
	}

	arg_list.PushBack(cl_kernel_arg)

	fmt.Printf("[ocl-wrapper] Set Kernel Arg, Kernel ID: %v, arg_idx: %v,  size: %v, value: %v, type: %v\n", kernel, arg_idx, size, cl_kernel_arg.ptr_val, cl_kernel_arg.arg_type)

	return 0 // CL_SUCCESS
}

// global_work_size is type uint32
// local_work_size is type uint16
//export gcn3LaunchKernel
func gcn3LaunchKernel(kernel int, global_work_size unsafe.Pointer, local_work_size unsafe.Pointer) int {
	var grid_args [3]uint32
	var work_args [3]uint16

	global := *(*[3]uint32)(global_work_size)
	local := *(*[3]uint16)(local_work_size)

	// Copy data over
	for i := 0; i < 3; i++ {
		grid_args[i] = global[i] //*(*uint32)(global_work_size) //uint32
		work_args[i] = local[i] //*(*uint16)(local_work_size)  //uint16
		fmt.Printf("[ocl-wrapper] Dimention: %i, global: %u, local: %u\n", i, grid_args[i], work_args[i])
	}

	cl_kernel := kernel_map[kernel]
	sim_kernel := cl_kernel.kernel
	arg_list := cl_kernel.arg_list

	// Reconcile kernel args
	num_args := arg_list.Len()
	args := make([]CLKernelArg, num_args)

	// Reorder args into slice
	for temp := arg_list.Front(); temp != nil; temp = temp.Next() {
		kernel_arg := temp.Value.(CLKernelArg)

		new_idx := kernel_arg.idx

		if new_idx >= num_args || new_idx < 0 {
			return -5 //FIXME add the right error
		}

		args[new_idx] = kernel_arg
	}

	/*
	   var all_args []byte
	   for _, kernel_arg := range args {
	       all_args = append(all_args, kernel_arg.bytes...)
	   }
	*/
	driverUpdateLDSPointers(sim_driver, sim_kernel, args)

	all_args := convertArgsToBytes(args)

	//kernel_arg_interface := createKernelArgInterface(args)

	sim_driver.LaunchKernelRuntimeArgs(sim_kernel, grid_args, work_args, all_args)

	return 1 // CL_SUCCESS
}

func main() {} // Required but ignored
