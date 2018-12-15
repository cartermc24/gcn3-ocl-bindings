# gcn3-ocl-bindings
OpenCL Bindings for the GCN3 Multi-GPU Simulator

## Building
### To build the front-end OpenCL library automatically:
1. Ensure the GCN3Sim simulator is installed (see getting started guide for GCN3)
2. Automatic: Run the build_ocl_frontend.sh script

### To build manually:
1. Ensure the GCN3Sim simulator is installed (see getting started guide for GCN3)
2. Enter the libsim-go directory and run make
3. Enter the libsim-ocl directory, create a build directory and cd into it
4. Run "cmake .." and then "make"
5. (If you are not running a ROCm system): Download the clang-ocl compiler package from https://mccardwell.net/simocl/simocl.tar.gz and untar the file in your home directory, it will create a .simocl folder with all the required files inside.
5. (If you are running a ROCm system): Create a .simocl folder in your home directory and simlink clang-ocl to your installation of clang-ocl inside the .simocl folder.
The frontend looks for ~/.simocl/clang-ocl when dynamically compiling OpenCL sources during runtime.

After building, there will be a libsim-ocl.so file created that contains both the OpenCL frontend as well as the GCN3Sim simulator.

## How to use the frontend:
To use the OpenCL frontend, you can either LD_PRELOAD the library to replace the native OpenCL library or if you are not running on an system that has OpenCL, you can build applications with the library directly.

### LD_PRELOAD:
- Create a simlink to libsim-ocl.so called libOpenCL.so and run "export LD_PRELOAD=/path/to/libOpenCL.so/simlink"
- The application will now use the simulator instead of the native OpenCL runtime
- To switch back to the native runtime, run "export LD_PRELOAD="

### Building an application with the simulator directly:
- When compiling an application, replace "-lOpenCL" with "-L/path/to/libsim_ocl.so -lsim_ocl"
- The application will now be linked with the simulator
- You might have to add the library to your LD_LIBRARY_PATH (run "export LD_LIBRARY_PATH=/path/to/libsim_ocl.so")
