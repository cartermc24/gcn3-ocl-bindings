package main

/*
 */
import "C"
import (
	"container/list"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"unsafe"
	"os"
	"os/user"
	"os/exec"
	"bytes"
//	"encoding/hex"
	"debug/elf"

	"gitlab.com/akita/gcn3/driver"
	"gitlab.com/akita/gcn3/insts"
	"gitlab.com/akita/gcn3/kernels"
	"gitlab.com/akita/gcn3/platform"
)

// Globals
var sim_driver *driver.Driver
var program_map map[int]*CLProgram
var buffer_map map[int]*driver.GPUPtr
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
/*
   build_status:
   	0: Not built yet
	1: Build failure
	2: Built successfully, program binary saved
	3: Released
*/
type CLProgram struct {
	program_string string
	program        []byte
	build_status   int
	build_status_msg string
}

type CLKernel struct {
	arg_list *list.List
	kernel   *insts.HsaCo
	kernel_name string
}

/*
   CLKernelArg Types
   0 -> GlobalPtr
   1 -> LocalPtr
   2 -> Primative
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
	for _, kernel_arg := range cl_kernel_args {
		arg_bytes := make([]byte, 8)//unsafe.Sizeof(kernel_arg.ptr_val))
		binary.LittleEndian.PutUint64(arg_bytes, uint64(kernel_arg.ptr_val)) // FIXME assumes 64-bit platform	
		//fmt.Printf("[ocl-wrapper] Argument %i has bytes: [%s]\n", i, hex.Dump(arg_bytes))
		all_args = append(all_args, arg_bytes...)
	}

	/*
	arg_bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(arg_bytes, 0)
	all_args = append(all_args, arg_bytes...)

	arg_bytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(arg_bytes, 400000)
	all_args = append(all_args, arg_bytes...)

	arg_bytes = make([]byte, 8)
	binary.LittleEndian.PutUint64(arg_bytes, 800000)
	all_args = append(all_args, arg_bytes...)

	arg_bytes = make([]byte, 4)
	binary.LittleEndian.PutUint32(arg_bytes, 100000)
	all_args = append(all_args, arg_bytes...)
	*/

	offsets := make([]byte, 192)
	all_args = append(all_args, offsets...)

	/*
	if len(all_args) % 64 != 0 {
		fmt.Printf("[ocl-wrapper] Kernel arguments are not aligned to 64-bits\n")
	}
	*/

	//fmt.Printf("[ocl-wrapper] Arguments passed to kernel [%s]", hex.Dump(all_args))

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
	//This function will return the number of simulated GPUs avaliable
	//the GPU ids will be the index from 1..(num gpus)
	return 1
}

//export gcn3GetKernelInfo
func gcn3GetKernelInfo(kernel int, param_name int, param_value_size uint64, param_ptr unsafe.Pointer, param_ptr_size unsafe.Pointer) int {
	var ptr_size uint64 = 100
	if param_ptr_size != nil {
		ptr_size = *(*uint64)(param_ptr_size)
	}

//	fmt.Printf("!!! KERNELID: %i !!!\n", kernel)

	switch param_name {
		case 0x1190: // CL_KERNEL_FUNCTION_NAME 
//			fmt.Printf("!!! QUERIED NAME IS: %s !!!\n", kernel_map[kernel].kernel_name)
			writeStringToPtr(param_ptr, ptr_size, kernel_map[kernel].kernel_name)
	}

	return 0 // CL_SUCCESS
}

//export gcn3GetProgramBuildInfo
func gcn3GetProgramBuildInfo(program int, device_id int, param_name int, param_value_size uint64, param_ptr unsafe.Pointer, param_ptr_size unsafe.Pointer) int {
	var ptr_size uint64 = 100
	if param_ptr_size != nil {
		ptr_size = *(*uint64)(param_ptr_size)
	}

	switch param_name {
		case 0x1181: // CL_PROGRAM_BUILD_STATUS
			fmt.Printf("Program [%i] has build status [%i]\n", program, program_map[program].build_status)
			if program_map[program].build_status == 0 {
				*(*int)(param_ptr) = -1 //CL_BUILD_NONE
			} else if program_map[program].build_status == 1 {
				*(*int)(param_ptr) = -2 //CL_BUILD_ERROR
			} else if program_map[program].build_status == 2 {
				*(*int)(param_ptr) = 0 //CL_BUILD_SUCCESS
			} else { // If we are in a weird state (like released) return CL_BUILD_NONE
				*(*int)(param_ptr) = -1
			}
		case 0x1183: // CL_PROGRAM_BUILD_LOG
			if program_map[program].build_status == 1 {
				writeStringToPtr(param_ptr, ptr_size, program_map[program].build_status_msg)
			} else {
				writeStringToPtr(param_ptr, ptr_size, "No error log")
			}
	}

	return 0 // CL_SUCCESS
}


//export gcn3GetDeviceInfo
func gcn3GetDeviceInfo(device_id int, param_name int, param_value_size uint64, param_ptr unsafe.Pointer, param_ptr_size unsafe.Pointer) int {
	var ptr_size uint64 = 1000
	if param_ptr_size != nil {
		ptr_size = *(*uint64)(param_ptr_size)
	}

	switch param_name {
		case 0x102B: // CL_DEVICE_NAME
			writeStringToPtr(param_ptr, ptr_size, "GCN3 Simulated GPU")
		case 0x102C: // CL_DEVICE_VENDOR
			writeStringToPtr(param_ptr, ptr_size, "NUCAR")
		case 0x1000: // CL_DEVICE_TYPE
			*(*uint)(param_ptr) = 4 // CL_DEVICE_TYPE_GPU
		case 0x1002: // CL_DEVICE_MAX_COMPUTE_UNITS
			*(*uint)(param_ptr) = 20
		case 0x1004: // CL_DEVICE_MAX_WORK_GROUP_SIZE
			*(*uint)(param_ptr) = 512//1024
		case 0x1023: // CL_DEVICE_LOCAL_MEM_SIZE
			*(*uint)(param_ptr) = 24576//49152
		case 0x1022: // CL_DEVICE_LOCAL_MEM_TYPE
			*(*uint)(param_ptr) = 0x1
		case 0x101F: // CL_DEVICE_GLOBAL_MEM_SIZE
			*(*uint)(param_ptr) = 1063837696//8510701568
		case 0x1027: // CL_DEVICE_AVALIABLE
			*(*uint)(param_ptr) = 1
		case 0x1053: // CL_DEVICE_SVM_CAPABILITIES
			*(*uint)(param_ptr) = 0 // Disable SVM
		case 0x1005: // CL_DEVICE_MAX_WORK_ITEM_SIZE
			max_wis := (*[3]C.size_t)(param_ptr)
			max_wis[0] = 256//1024
			max_wis[1] = 256//1024
			max_wis[2] = 16//64
		case 0x1003: // CL_DEVICE_MAX_WORK_ITEM_DIMENSIONS
			*(*uint)(param_ptr) = 3 // Always 3D
		case 0x1001: // CL_DEVICE_VENDOR_ID
			*(*uint)(param_ptr) = 0x1002 // Randomly chose 827 as vendor id
		case 0x1030: // CL_DEVICE_EXTENSIONS
			supported_ext := "cl_khr_global_int32_base_atomics cl_khr_global_int32_extended_atomics cl_khr_local_int32_base_atomics cl_khr_local_int32_extended_atomics cl_khr_fp64 cl_khr_byte_addressable_store cl_khr_icd cl_khr_gl_sharing"
			writeStringToPtr(param_ptr, ptr_size, supported_ext)
		case 0x1035: // CL_DEVICE_HOST_UNIFIED_MEMORY
			*(*uint)(param_ptr) = 0 // No host unified memory
	}

	return 0 // CL_SUCCESS
}

func writeStringToPtr(ptr unsafe.Pointer, ptr_len uint64, str string) {
	var strptr = uintptr(ptr)
	var str_len uint64
	if uint64(len(str)) > ptr_len {
		str_len = ptr_len
	} else {
		str_len = uint64(len(str))
	}

	for i := uint64(0); i < str_len; i++ {
                *(*C.uchar)(unsafe.Pointer(strptr)) = C.uchar(str[i])
                strptr++
        }

	*(*C.uchar)(unsafe.Pointer(strptr)) = C.uchar(0)
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
	cl_program.build_status = 0

	program_map[program_idx] = &cl_program

	program_idx += 1

//	fmt.Printf("[ocl-wrapper] Created program with ID %v\n", program_idx-1)

	return program_idx - 1
}

//export gcn3BuildProgram
func gcn3BuildProgram(program_id int) int {
	// Set build status to failure, we will reset if successful
	program_map[program_id].build_status = 1

	// Write program source to file
//	fmt.Printf("[ocl-wrapper] Writing CL source to temporary file\n")
	program_bytes := []byte(program_map[program_id].program_string)
	write_err := ioutil.WriteFile("/tmp/prog.cl", program_bytes, 0644)
	if write_err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: could not write CL source to /tmp\n")
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

	// Run compiler
//	fmt.Printf("[ocl-wrapper] Running clang-ocl\n")
	usr, usr_err := user.Current()
	if usr_err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: unable to get user information to locate compiler\n")
		program_map[program_id].build_status_msg = "[ocl-wrapper] Error: unable to get user information to locate compiler"
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
		program_map[program_id].build_status_msg = stderr.String()
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ocl-wrapper] Error: could not invoke compiler, ensure clang-ocl exists at ~/.clangocl/clang-ocl, error: %s\n", err)
		program_map[program_id].build_status_msg = "[ocl-wrapper] Error: could not invoke compiler, ensure clang-ocl exists at ~/.clangocl/clang-ocl"
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

//	fmt.Printf("[ocl-wrapper] CL source successfully compiled\n")
	hsacoBytes, err := ioutil.ReadFile("/tmp/prog.hsaco")
	if err != nil {
		fmt.Printf("[ocl-wrapper] Error building program: %v\n", err)
		program_map[program_id].build_status_msg = "[ocl-wrapper] Error building program, no output from compiler"
		return -11 // CL_BUILD_PROGRAM_FAILURE
	}

	program_map[program_id].program = hsacoBytes
	program_map[program_id].build_status = 2 // Build success

//	fmt.Printf("[ocl-wrapper] Built program with ID %v\n", program_id)
	return 0 // CL_SUCCESS
}

//export gcn3CreateKernelsInProgram
func gcn3CreateKernelsInProgram(program_id int, num_kernels uint, kernels unsafe.Pointer, num_kernels_ret unsafe.Pointer, cl_kernel_size int) int {
	reader := bytes.NewReader(program_map[program_id].program)
	executable, err := elf.NewFile(reader)
	if err != nil {
		fmt.Printf("[ocl-wrapper]: Error %s\n", err)
		return -47 // CL_INVALID_KERNEL_DEFINITION
	}

	symbols, err := executable.DynamicSymbols()
	if err != nil {
		fmt.Printf("[ocl-wrapper]: Error %s\n", err)
		return -47 // CL_INVALID_KERNEL_DEFINITION
	}

	var kernelptr = uintptr(kernels)
	var i = uint(0)
	for _, symbol := range symbols {
		*(*C.int)(unsafe.Pointer(kernelptr)) = C.int(gcn3CreateKernel(program_id, symbol.Name))
//		fmt.Printf("[ocl-wrapper] Created kernel with name: [%s]\n", symbol.Name)
		kernelptr += uintptr(cl_kernel_size)
		if i >= num_kernels {
			break
		}
		i++
	}

	if num_kernels_ret != nil {
		*(*uint64)(num_kernels_ret) = uint64(i)
	}

	return 0 // CL_SUCCESS
}

//export gcn3CreateKernel
func gcn3CreateKernel(program_id int, kernel_name string) int {

	//fmt.Println("[ocl-wrapper] Attempting to create kernel")
	program := program_map[program_id].program

	cl_kernel := CLKernel{}
	kernel := kernels.LoadProgramFromMemory(program, kernel_name)

	if kernel == nil {
		fmt.Println("[ocl-wrapper] Unable to get kernel from program")
		return -46 // CL_INVALID_KERNEL_NAME
	}

	cl_kernel.kernel = kernel
	cl_kernel.arg_list = list.New()
	cl_kernel.kernel_name = kernel_name
	kernel_map[kernel_idx] = &cl_kernel

	kernel_idx += 1

	//fmt.Printf("[ocl-wrapper] Created kernel with name: %v, ID: %v\n", kernel_name, kernel_idx-1)
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

//	fmt.Printf("[ocl-wrapper] Allocated buffer of size: %v, wrapper ID: %v, sim addr: %v\n", size, buffer_idx-1, new_buffer)
	return buffer_idx - 1
}

//export gcn3EnqueueWriteBuffer
func gcn3EnqueueWriteBuffer(buffer int, size int, ptr unsafe.Pointer) int {
	sim_buffer := buffer_map[buffer]

	ptr_bytes := C.GoBytes(ptr, C.int(size))

	sim_driver.MemoryCopyHostToDevice(*sim_buffer, ptr_bytes)

//	fmt.Printf("[ocl-wrapper] Wrote data to device: [%s] @ region: %02x\n", hex.Dump(back), *sim_buffer)

	//fmt.Printf("[ocl-wrapper] Enqueued Write Buffer for buffer ID %v\n", buffer)
	return 0 // CL_SUCCESS
}

//export gcn3EnqueueReadBuffer
func gcn3EnqueueReadBuffer(buffer int, size int, ptr unsafe.Pointer) int {
	sim_buffer := buffer_map[buffer]

	ptr_bytes := make([]byte, size)
	for i := 0; i < size; i++ {
		ptr_bytes[i] = 0xFF
	}

	sim_driver.MemoryCopyDeviceToHost(ptr_bytes, *sim_buffer)

//	fmt.Printf("[ocl-wrapper] Fetched data from device: [%s] @ region %02x\n", hex.Dump(ptr_bytes), *sim_buffer)

	var cptr = uintptr(ptr)
	for i := 0; i < size; i++ {
		*(*C.uchar)(unsafe.Pointer(cptr)) = C.uchar(ptr_bytes[i])
		cptr++
	}

	//fmt.Printf("[ocl-wrapper] Enqueued Read Buffer for buffer ID %v\n", buffer)

	return 0 // CL_SUCCESS
}

//export gcn3SetKernelArg
func gcn3SetKernelArg(kernel int, arg_idx int, size int, ptr unsafe.Pointer) int {
	cl_kernel := kernel_map[kernel]
	arg_list := cl_kernel.arg_list
	cl_kernel_arg := CLKernelArg{}
	cl_kernel_arg.idx = arg_idx
	cl_kernel_arg.size = size

	if ptr == nil {
//		fmt.Printf("[ocl-wrapper] Arg %i is a LOCAL\n", arg_idx)
		cl_kernel_arg.ptr_val = 0
		cl_kernel_arg.arg_type = 1 // Local
	} else {
		ptr_value := uint64(*(*uint32)(ptr))
//		fmt.Printf("[ocl-wrapper] Input PTR is: %i\n", ptr_value)
		if val, ok := buffer_map[(int)(ptr_value)]; ok {
//			fmt.Printf("[ocl-wrapper] Arg %i is a GLOBAL\n", arg_idx)
			cl_kernel_arg.ptr_val = (uint64)(*val)
			cl_kernel_arg.arg_type = 0 // Global
		} else {
//			fmt.Printf("[ocl-wrapper] Arg %i is a PRIMATIVE\n", arg_idx)
			cl_kernel_arg.ptr_val = ptr_value
			cl_kernel_arg.arg_type = 2 // Primative
		}
	}

/*
	test := unsafe.Sizeof(reflect.TypeOf((int)(0)))

	if uintptr(size) > test {
		cl_kernel_arg.arg_type = uint8(1)
	} else {
		cl_kernel_arg.arg_type = uint8(0)
	}
*/

	arg_list.PushBack(cl_kernel_arg)

	//fmt.Printf("[ocl-wrapper] Set Kernel Arg, Kernel ID: %v, arg_idx: %v,  size: %v, value: %v, type: %v\n", kernel, arg_idx, size, cl_kernel_arg.ptr_val, cl_kernel_arg.arg_type)

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
//		fmt.Printf("[ocl-wrapper] Dimention: %i, global: %u, local: %u\n", i, grid_args[i], work_args[i])
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

	return 0 // CL_SUCCESS
}

//export gcn3ReleaseProgram
func gcn3ReleaseProgram(program_id int) int {
	program_map[program_id].program = nil
	program_map[program_id].program_string = ""
	program_map[program_id].build_status = 3
	return 0 // CL_SUCCESS
}

func main() {} // Required but ignored
