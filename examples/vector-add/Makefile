all: checkEnv clean vecAdd_cl vecAdd_cc_cl

.PHONY: checkEnv clean

vecAdd_cl: vecAdd.c
	mkdir -p bin
	cc -o bin/vecAdd_cl vecAdd.c -L../../libsim-ocl/build -lsim_ocl -lm -lpthread -g

vecAdd_native_cl: vecAdd.cc
	mkdir -p bin
	CC -lOpenCL -o bin/vecAdd_ccl vecAdd.cc

clean:
	rm -rf bin
