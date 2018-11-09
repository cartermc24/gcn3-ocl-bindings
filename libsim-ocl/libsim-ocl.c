#include <string.h>
#include "libsim-ocl.h"

void initialize_simulator() {
    if (is_simulator_initialized) {
        return;
    }
    is_simulator_initialized = true;

    initializeSimulator();
}

cl_int clGetPlatformIDs(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms) {
    initialize_simulator(); 
    return (cl_int)gcn3GetPlatformIDs();
}

cl_int clGetDeviceIDs(cl_platform_id platform,
                      cl_device_type device_type,
                      cl_uint num_entries,
                      cl_device_id *devices,
                      cl_uint *num_devices) {
    initialize_simulator(); 
    return (cl_int)gcn3GetDeviceIDs();
}

cl_int clGetDeviceInfo(cl_device_id device,
                       cl_device_info param_name,
                       size_t param_value_size,
                       void *param_value,
                       size_t *param_value_size_ret) {
    initialize_simulator(); 
    return CL_SUCCESS;
}

cl_context clCreateContext(const cl_context_properties *properties,
                cl_uint num_devices,
                const cl_device_id *devices,
                void (CL_CALLBACK *pfn_notify)(const char *, const void *, size_t, void *),
                void *user_data,
                cl_int *errcode_ret) {
    *errcode_ret = CL_SUCCESS;
    initialize_simulator(); 
    return (cl_context)gcn3CreateContext();
}

cl_command_queue clCreateCommandQueue(cl_context context,
                                      cl_device_id device,
                                      cl_command_queue_properties properties,
                                      cl_int *errcode_ret) {
    initialize_simulator(); 
    *errcode_ret = CL_SUCCESS;
    cl_command_queue queue;
    return queue;
}

#ifndef __APPLE__
cl_command_queue clCreateCommandQueueWithProperties(cl_context context,
                                                    cl_device_id device,
                                                    const cl_queue_properties *properties,
                                                    cl_int *errcode_ret) {
    initialize_simulator(); 
    cl_command_queue queue;
    return queue;
}
#endif

cl_program clCreateProgramWithSource(cl_context context,
                                     cl_uint count,
                                     const char **strings,
                                     const size_t *lengths,
                                     cl_int *errcode_ret) {
    initialize_simulator(); 
    GoString string;
    string.p = strings[0];
    string.n = lengths[0];

    int response = (int)gcn3CreateProgramWithSource((GoInt)context, string);
    if (response < 0) {
        *errcode_ret = response;
        return NULL;
    } else {
        *errcode_ret = CL_SUCCESS;
        return response;
    }

}

cl_int clBuildProgram(cl_program program,
                      cl_uint num_devices,
                      const cl_device_id *device_list,
                      const char *options,
                      void (*pfn_notify)(cl_program, void *user_data),
                      void *user_data) {
    initialize_simulator(); 

    return (cl_int)gcn3BuildProgram((GoInt)program);
}

cl_int clGetProgramBuildInfo(cl_program program,
                             cl_device_id device,
                             cl_program_build_info param_name,
                             size_t param_value_size,
                             void *param_value,
                             size_t *param_value_size_ret) {
    initialize_simulator(); 
    return CL_SUCCESS;
}

cl_mem clCreateBuffer(cl_context context,
                      cl_mem_flags flags,
                      size_t size,
                      void *host_ptr,
                      cl_int *errcode_ret) {
    initialize_simulator(); 

    int buffer_desc = (int)gcn3CreateBuffer((GoInt)context, (GoInt)size);
    if (buffer_desc < 0) {
        *errcode_ret = buffer_desc;
        return NULL;
    } else {
        *errcode_ret = CL_SUCCESS;
        return buffer_desc;
    }
}

cl_kernel clCreateKernel(cl_program program,
                         const char *kernel_name,
                         cl_int *errcode_ret) {
    initialize_simulator(); 

    GoString string;
    string.p = kernel_name;
    string.n = strlen(kernel_name);

    int kernel_desc = (int)gcn3CreateKernel((GoInt)program, string);

    if (kernel_desc < 0) {
        *errcode_ret = kernel_desc;
        return NULL;
    } else {
        *errcode_ret = CL_SUCCESS;
        return kernel_desc;
    }
}

cl_int clSetKernelArg(cl_kernel kernel,
                      cl_uint arg_index,
                      size_t arg_size,
                      const void *arg_value) {
    initialize_simulator();
    return (cl_int)gcn3SetKernelArg((GoInt)kernel, (GoInt)arg_index, (GoInt)arg_size, arg_value);
}

cl_int clEnqueueNDRangeKernel(cl_command_queue command_queue,
                              cl_kernel kernel,
                              cl_uint work_dim,
                              const size_t *global_work_offset,
                              const size_t *global_work_size,
                              const size_t *local_work_size,
                              cl_uint num_events_in_wait_list,
                              const cl_event *event_wait_list,
                              cl_event *event) {
    initialize_simulator(); 

    uint32_t l_global_work_size[3] = { 1 };
    uint16_t l_local_work_size[3] = { 1 };

    for (int i = 0; i < work_dim; i++) {
        l_global_work_size[i] = (uint32_t)global_work_size[i];
        l_local_work_size[i] = (uint16_t)local_work_size[i];
    }

    return (cl_int)gcn3LaunchKernel((GoInt)kernel, &l_global_work_size, &l_local_work_size);
}

cl_int clWaitForEvents(cl_uint num_events,
                       const cl_event *event_list) {
    initialize_simulator(); 
    return CL_SUCCESS;
}

cl_int clEnqueueWriteBuffer(cl_command_queue command_queue,
                            cl_mem buffer,
                            cl_bool blocking_read,
                            size_t offset,
                            size_t size,
                            const void *ptr,
                            cl_uint num_events_in_wait_list,
                            const cl_event *event_wait_list,
                            cl_event *event) {
    initialize_simulator(); 
    return (cl_int)gcn3EnqueueWriteBuffer((GoInt)buffer, (GoInt)size, ptr);
}

cl_int clEnqueueReadBuffer(cl_command_queue command_queue,
                           cl_mem buffer,
                           cl_bool blocking_read,
                           size_t offset,
                           size_t size,
                           void *ptr,
                           cl_uint num_events_in_wait_list,
                           const cl_event *event_wait_list,
                           cl_event *event) {
    initialize_simulator(); 
    return (cl_int)gcn3EnqueueReadBuffer((GoInt)buffer, (GoInt)size, ptr);
}

cl_int clReleaseMemObject(cl_mem memobj) {
    initialize_simulator(); 
    return CL_SUCCESS;
}

cl_int clReleaseContext(cl_context context) {
    initialize_simulator(); 
    return CL_SUCCESS;
}

cl_int clReleaseCommandQueue(cl_command_queue command_queue) {
    initialize_simulator(); 
    return CL_SUCCESS;
}
