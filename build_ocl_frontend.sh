#!/bin/bash

if ! [ -x "$(command -v wget)" ]; then
  echo 'Error: wget is not installed and is required.' >&2
  exit 1
fi

if ! [ -x "$(command -v cmake)" ]; then
  echo 'Error: cmake is not installed and is required.' >&2
  exit 1
fi

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
echo "[COMPILER]: Downloading clang-ocl"
wget https://mccardwell.net/simocl/simocl.tar.gz
tar xzvf simocl.tar.gz
echo "[COMPILER]: Moving clang-ocl to ~/.simocl"
mv home/ubuntu/.simocl ~/.simocl
rm -rf home/ simocl.tar.gz
echo "[COMPILER]: Done"
echo "[TEST]: Setting LD_LIBRARY_PATH"
export LD_LIBRARY_PATH=$(pwd)/libsim-ocl/build:$LD_LIBRARY_PATH
echo "[TEST]: Running vector-add test"
cd examples/vector-add/
make vecAdd_cl
./bin/vecAdd_cl
cd ../../
echo "[TEST]: Done, ensure output is 100"
