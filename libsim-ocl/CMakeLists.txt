cmake_minimum_required(VERSION 3.12)
project(libsim_ocl C)

set(CMAKE_C_STANDARD 99)

find_package( OpenCL REQUIRED )
include_directories( ${OpenCL_INCLUDE_DIR} )

add_library(sim_ocl SHARED libsim-ocl.c libsim-ocl.h)
