cmake_minimum_required(VERSION 3.9.5)
project(libsim_ocl C)

set(CMAKE_C_STANDARD 99)

find_package( OpenCL REQUIRED )
include_directories( ${OpenCL_INCLUDE_DIR} )
include_directories( ../libsim-go/_obj )

add_library(sim_ocl SHARED libsim-ocl.c libsim-ocl.h)
target_link_libraries(sim_ocl ${CMAKE_CURRENT_SOURCE_DIR}/../libsim-go/_obj/_cgo_.o)