all: checkEnv clean vecAdd_cl vecAdd_cc_cl

.PHONY: checkEnv clean

vecAdd_cl: vecAdd.c
	mkdir -p bin
	cc -o bin/vecAdd_cl vecAdd.c -L/home/shance/Documents/gpu-research/gcn3-ocl-bindings/libsim-ocl/build -lsim_ocl -lm -lpthread

vecAdd_cl_ref: vecAdd.c
	mkdir -p bin
	cc -o bin/vecAdd_cl_ref vecAdd.c -L/opt/amdgpu-pro/lib/x86_64-linux-gnu -lOpenCL -lm

vecAdd_cc_cl: vecAdd.cc
	mkdir -p bin
	CC -lOpenCL -o bin/vecAdd_ccl vecAdd.cc

checkEnv:
ifndef CRAY_CUDATOOLKIT_VERSION
	$(error cudatoolkit module not loaded)
endif

clean:
	rm -rf bin
