package main

/*
*/
import "C"
import (
    "unsafe"
    "container/list"
    "io/ioutil"
    "reflect"

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
type CLProgram struct {
    program_string string
    program []byte
}

type CLKernel struct {
    arg_list *list.List
    kernel *insts.HsaCo
}

/*
    CLKernelArg Types
    0 -> Global Ptr
    1 -> Local Ptr
    2 -> Primitive
*/
type CLKernelArg struct {
    idx int
    size int
    ptr_val unsafe.Pointer
    arg_type int
}


// 
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
}


//export initializeSimulator
func initializeSimulator() {
    buffer_map = make(map[int]*driver.GPUPtr)
    kernel_map = make(map[int]*CLKernel)
    program_map = make(map[int]*CLProgram)
    _, _,sim_driver,_ = platform.BuildEmuPlatform()
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
    if (context != context_id) {
        return -34 // CL_INVALID_CONTEXT
    }

    cl_program := CLProgram{}
    cl_program.program_string = program_string

    program_map[program_idx] = &cl_program

    program_idx += 1

    return program_idx - 1
}


//export gcn3BuildProgram
func gcn3BuildProgram(program_id int) int {
    // FIXME actually build program
    hsacoBytes, err := ioutil.ReadFile("/home/shance/Documents/gpu-research/myfirstkernel.hsaco")
    if (err != nil) {
        return -11 // CL_BUILD_PROGRAM_FAILURE
    }

    program_map[program_id].program = hsacoBytes

    return 1 // CL_SUCCESS
}


//export gcn3CreateKernel
func gcn3CreateKernel(program_id int, kernel_name string) int {
    program := program_map[program_id].program

    cl_kernel := CLKernel{}
    kernel := kernels.LoadProgramFromMemory(program, kernel_name)
    if (kernel == nil) {
        return -46 // CL_INVALID_KERNEL_NAME
    }

    cl_kernel.kernel = kernel
    cl_kernel.arg_list = list.New()
    kernel_map[kernel_idx] = &cl_kernel

    kernel_idx += 1

    return kernel_idx - 1
}


// Returns Buffer ID
// Size in bytes
//export gcn3CreateBuffer
func gcn3CreateBuffer(context int, size int) int {
    if (context != context_id) {
        return -34 // CL_INVALID_CONTEXT
    }

    new_buffer := sim_driver.AllocateMemory(uint64(size))
    buffer_map[buffer_idx] = &new_buffer

    buffer_idx += 1
    return buffer_idx - 1
}


//export gcn3EnqueueWriteBuffer
func gcn3EnqueueWriteBuffer(buffer int, size int, ptr unsafe.Pointer) int {
    sim_buffer := buffer_map[buffer]

    ptr_bytes := C.GoBytes(ptr, C.int(size))

    sim_driver.MemoryCopyHostToDevice(*sim_buffer, ptr_bytes)

    //ptr_bytes = *((*[size]byte) (ptr))

    //num_bytes_copied := copy(sim_buffer, ptr_bytes)

    //if (num_bytes_copied != size) {
    //    return -37 // CL_INVALID_HOST_PTR
    //}

    return 0 // CL_SUCCESS
}


//export gcn3EnqueueReadBuffer
func gcn3EnqueueReadBuffer(buffer int, size int, ptr unsafe.Pointer) int {
    sim_buffer := buffer_map[buffer]

    ptr_bytes := C.GoBytes(ptr, C.int(size))

    sim_driver.MemoryCopyDeviceToHost(ptr_bytes, *sim_buffer)

    //ptr_bytes := *((*[size]byte) (ptr))

    /*
    num_bytes_copied := copy(ptr_bytes[:], sim_buffer[:])

    if (num_bytes_copied != size) {
        return -37 // CL_INVALID_HOST_PTR
    }
    */

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
    cl_kernel_arg.ptr_val = ptr

    test := unsafe.Sizeof(reflect.TypeOf((int) (0)))

    if uintptr(size) > test {
        cl_kernel_arg.arg_type = 1
    } else {
        cl_kernel_arg.arg_type = 0
    }

    arg_list.PushBack(cl_kernel_arg)

    return 0 // CL_SUCCESS
}


// global_work_size is type uint32
// local_work_size is type uint16
//export gcn3LaunchKernel
func gcn3LaunchKernel(kernel int, global_work_size unsafe.Pointer, local_work_size unsafe.Pointer) int {
    var grid_args [3]uint32
    var work_args [3]uint16

    // Copy data over
    for i := 0; i < 3; i++ {
        grid_args[i] = *(*uint32) (unsafe.Pointer(uintptr(global_work_size) + uintptr(4))) //uint32
        work_args[i] = *(*uint16) (unsafe.Pointer(uintptr(local_work_size) + uintptr(2))) //uint16
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

        if (new_idx >= num_args || new_idx < 0) {
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
    kernel_arg_interface := createKernelArgInterface(args)

    sim_driver.LaunchKernel(sim_kernel, grid_args, work_args, kernel_arg_interface)

    return 1 // CL_SUCCESS
}


func main() {} // Required but ignored
