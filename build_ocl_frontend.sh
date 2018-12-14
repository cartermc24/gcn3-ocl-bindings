#!/bin/bash

echo "[BUILD]: Building GCN3 OpenCL frontend"
echo "[BUILD]: Generating CGo C headers"
cd libsim-go
make
cd ..
echo "[BUILD]: Building C wrapper"
cd libsim-ocl
rm -rf build/
mkdir build
cd build
cmake ..
make
cd ../../
echo "[BUILD]: Done"
echo "[TEST]: Setting LD_LIBRARY_PATH"
export LD_LIBRARY_PATH=$(pwd)/libsim-ocl/build:$LD_LIBRARY_PATH
echo "[TEST]: Running vector-add test"
cd examples/vector-add/
make vecAdd_cl
./bin/vecAdd_cl
cd ../../
echo "[TEST]: Done, ensure output is 100"
