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
#ifdef TRACE
    printf("[OCL-TRACE]: clGetPlatformIDs called with [%u, %p, %u]\n", num_entries, platforms, *num_platforms);
#endif

    return (cl_int)gcn3GetPlatformIDs();
}

cl_int clGetDeviceIDs(cl_platform_id platform,
                      cl_device_type device_type,
                      cl_uint num_entries,
                      cl_device_id *devices,
                      cl_uint *num_devices) {
    initialize_simulator(); 
#ifdef TRACE
    printf("[OCL-TRACE]: clGetDeviceIDs called with device type %i\n", device_type);
#endif

    if (device_type != CL_DEVICE_TYPE_ALL || device_type != CL_DEVICE_TYPE_GPU || device_type != CL_DEVICE_TYPE_DEFAULT) {
	fprintf(stderr, "[ocl-wrapper]: Warning: Attempted to query non-GPU devices...\n");
    }

    if (devices == NULL) {
	fprintf(stderr, "[ocl-wrapper]: Warning: devices is null, no devices can be returned\n");
	return CL_SUCCESS;
    }

    int32_t num_sim_devices = gcn3GetDeviceIDs();
    if (num_devices != NULL) { *num_devices = num_sim_devices; }

    int32_t max_entries = num_sim_devices > num_entries ? num_entries : num_sim_devices;
    for (int32_t i = 0; i < max_entries; i++) {
	devices[0] = i+1;
    }

    return CL_SUCCESS;
}

cl_int clGetDeviceInfo(cl_device_id device,
                       cl_device_info param_name,
                       size_t param_value_size,
                       void *param_value,
                       size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetDeviceInfo called asking for %04x\n", param_name);
#endif

    return (cl_int)gcn3GetDeviceInfo(device, param_name, param_value_size, param_value, param_value_size_ret);
}

cl_int clGetContextInfo(cl_context context,
  			cl_context_info param_name,
  			size_t param_value_size,
  			void *param_value,
  			size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetContextInfo called\n");
#endif

    return CL_SUCCESS;
}

cl_context clCreateContext(const cl_context_properties *properties,
                cl_uint num_devices,
                const cl_device_id *devices,
                void (CL_CALLBACK *pfn_notify)(const char *, const void *, size_t, void *),
                void *user_data,
                cl_int *errcode_ret) {
    
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateContext called\n");
#endif

    if (errcode_ret != NULL) { *errcode_ret = CL_SUCCESS; }
    return (cl_context)gcn3CreateContext();
}

cl_command_queue clCreateCommandQueue(cl_context context,
                                      cl_device_id device,
                                      cl_command_queue_properties properties,
                                      cl_int *errcode_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateCommandQueue called\n");
#endif

    if (errcode_ret != NULL) { *errcode_ret = CL_SUCCESS; }
    cl_command_queue queue;
    return queue;
}

#ifndef __APPLE__
cl_command_queue clCreateCommandQueueWithProperties(cl_context context,
                                                    cl_device_id device,
                                                    const cl_queue_properties *properties,
                                                    cl_int *errcode_ret) {
    initialize_simulator(); 
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateCommandQueueWithProperties called\n");
#endif

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
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateProgramWithSource called\n");
#endif

    GoString string;
    string.p = strings[0];

    if (lengths != NULL) {
        string.n = lengths[0];
    } else {
	string.n = strlen(strings[0]);
    }

    //printf("[ocl-wrapper-C]: Got %u source strings with len %i\n", count, string.n);

    int response = (int)gcn3CreateProgramWithSource((GoInt)context, string);
    if (response < 0) {
        if (errcode_ret != NULL) { *errcode_ret = response; }
        return NULL;
    } else {
        if (errcode_ret != NULL) { *errcode_ret = CL_SUCCESS; }
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
#ifdef TRACE
    printf("[OCL-TRACE]: clBuildProgram called\n");
#endif


    return (cl_int)gcn3BuildProgram((GoInt)program);
}

cl_int clGetProgramBuildInfo(cl_program program,
                             cl_device_id device,
                             cl_program_build_info param_name,
                             size_t param_value_size,
                             void *param_value,
                             size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetProgramBuildInfo called for program %u asking for param: %u\n", program, param_name);
#endif

    return (cl_int)gcn3GetProgramBuildInfo(program, device, param_name, param_value_size, param_value, param_value_size_ret);
}

cl_int clGetPlatformInfo(cl_platform_id platform,
  			 cl_platform_info param_name,
  			 size_t param_value_size,
  			 void *param_value,
  			 size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetPlatformInfo called\n");
#endif

    return CL_SUCCESS;
}

cl_mem clCreateBuffer(cl_context context,
                      cl_mem_flags flags,
                      size_t size,
                      void *host_ptr,
                      cl_int *errcode_ret) {
    initialize_simulator(); 
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateBuffer called\n");
#endif

    int buffer_desc = (int)gcn3CreateBuffer((GoInt)context, (GoInt)size);
    if (buffer_desc < 0) {
        if (errcode_ret != NULL) { *errcode_ret = buffer_desc; }
        return NULL;
    } else {
        if (errcode_ret != NULL) { *errcode_ret = CL_SUCCESS; }
        return buffer_desc;
    }
}

cl_kernel clCreateKernel(cl_program program,
                         const char *kernel_name,
                         cl_int *errcode_ret) {
    initialize_simulator(); 
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateKernel called\n");
#endif

    GoString string;
    string.p = kernel_name;
    string.n = strlen(kernel_name);

    int kernel_desc = (int)gcn3CreateKernel((GoInt)program, string);

    if (kernel_desc < 0) {
        if (errcode_ret != NULL) { *errcode_ret = kernel_desc; }
        return NULL;
    } else {
        if (errcode_ret != NULL) { *errcode_ret = CL_SUCCESS; }
        return kernel_desc;
    }
}

cl_int clSetKernelArg(cl_kernel kernel,
                      cl_uint arg_index,
                      size_t arg_size,
                      const void *arg_value) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clSetKernelArg called\n");
#endif

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
#ifdef TRACE
    printf("[OCL-TRACE]: clEnqueueNDRangeKernel called\n");
#endif

    uint32_t l_global_work_size[3];
    uint16_t l_local_work_size[3];

    for (int i = 0; i < 3; i++) {
	    l_global_work_size[i] = 1;
	    l_local_work_size[i] = 1;
    }

    for (int i = 0; i < work_dim; i++) {
        l_global_work_size[i] = (uint32_t)global_work_size[i];
        l_local_work_size[i] = (uint16_t)local_work_size[i];
    }

    return (cl_int)gcn3LaunchKernel((GoInt)kernel, l_global_work_size, l_local_work_size);
}

cl_int clWaitForEvents(cl_uint num_events,
                       const cl_event *event_list) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clWaitForEvents called\n");
#endif

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
#ifdef TRACE
    printf("[OCL-TRACE]: clEnqueueWriteBuffer called\n");
#endif

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
#ifdef TRACE
    printf("[OCL-TRACE]: clEnqueueReadBuffer called\n");
#endif

    return (cl_int)gcn3EnqueueReadBuffer((GoInt)buffer, (GoInt)size, ptr);
}

cl_int clEnqueueCopyBuffer(cl_command_queue command_queue,
  			   cl_mem src_buffer,
  			   cl_mem dst_buffer,
  			   size_t src_offset,
  			   size_t dst_offset,
  			   size_t cb,
  			   cl_uint num_events_in_wait_list,
  			   const cl_event *event_wait_list,
  			   cl_event *event) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clEnqueueCopyBuffer called\n");
#endif

    return CL_SUCCESS;
}

cl_int clFinish(cl_command_queue command_queue) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clFinish called\n");
#endif

    return CL_SUCCESS;
}

cl_int clRetainContext(cl_context context) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clRetainContext called\n");
#endif

    return CL_SUCCESS;
}

cl_int clReleaseMemObject(cl_mem memobj) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clReleaseMemObject called\n");
#endif
 
    return CL_SUCCESS;
}

cl_int clReleaseContext(cl_context context) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clReleaseContext called\n");
#endif

    return CL_SUCCESS;
}

cl_int clReleaseCommandQueue(cl_command_queue command_queue) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clReleaseCommandQueue called\n");
#endif

    return CL_SUCCESS;
}

cl_program clCreateProgramWithBinary(cl_context context,
				     cl_uint num_devices,
				     const cl_device_id *device_list,
				     const size_t *lengths,
				     const unsigned char **binaries,
				     cl_int *binary_status,
				     cl_int *errcode_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateProgramWithBinary called\n");
#endif

    return 0;
}

cl_int clGetProgramInfo(cl_program program,
		        cl_program_info param_name,
			size_t param_value_size,
			void *param_value,
			size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetProgramInfo called\n");
#endif

    return CL_SUCCESS;
}

cl_int clCreateKernelsInProgram(cl_program program,
				cl_uint num_kernels,
				cl_kernel *kernels,
				cl_uint *num_kernels_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clCreateKernelsInProgram called\n");
#endif
    
    memset(kernels, 0, num_kernels*sizeof(cl_kernel));

    return (cl_int)gcn3CreateKernelsInProgram(program, num_kernels, kernels, num_kernels_ret, sizeof(cl_kernel));
}

cl_int clGetKernelInfo(cl_kernel kernel,
  		       cl_kernel_info param_name,
  		       size_t param_value_size,
  	               void *param_value,
  		       size_t *param_value_size_ret) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clGetKernelInfo called with param: %u\n", param_name);
#endif

    return (cl_int)gcn3GetKernelInfo(kernel, param_name, param_value_size, param_value, param_value_size_ret);
}

cl_int clEnqueueTask(cl_command_queue command_queue,
  		     cl_kernel kernel,
  		     cl_uint num_events_in_wait_list,
  		     const cl_event *event_wait_list,
  		     cl_event *event) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clEnqueueTask called\n");
#endif

    return CL_SUCCESS;
}

cl_int clRetainCommandQueue(cl_command_queue command_queue) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clRetainCommandQueue called\n");
#endif

    return CL_SUCCESS;
}

cl_int clRetainMemObject(cl_mem memobj) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clRetainMemObject called\n");
#endif

    return CL_SUCCESS;
}

cl_int clReleaseProgram(cl_program program) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clReleaseProgram called\n");
#endif

    return (cl_int)gcn3ReleaseProgram(program);
}

cl_int clReleaseKernel(cl_kernel kernel) {
    initialize_simulator();
#ifdef TRACE
    printf("[OCL-TRACE]: clReleaseKernel called\n");
#endif

    return CL_SUCCESS;
}












































