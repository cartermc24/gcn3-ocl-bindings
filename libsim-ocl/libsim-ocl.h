//
// Created by Carter McCardwell on 9/30/18.
//

#ifndef LIBSIM_OCL_LIBSIM_OCL_H
#define LIBSIM_OCL_LIBSIM_OCL_H

#ifdef __APPLE__
#include <OpenCL/cl.h>
#else
#include <CL/cl.h>
#endif

#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#include "_cgo_export.h"

// ------TYPES------------
struct gpu_detail {
    uint64_t total_memory;
    uint32_t max_compute_units;
    uint32_t kernel_workgroup_size;
};

struct simulator_context {
    uint8_t num_gpus;
};

// ------MEMORY ELEMENTS---------------
bool is_simulator_initialized = false;
struct simulator_context context;
struct gpu_detail *gpu_details;

// ------FUNCTIONS---------------------
void initialize_simulator();

#endif //LIBSIM_OCL_LIBSIM_OCL_H
