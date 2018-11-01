all:
	gcc main_ocl12.c -o main -L../gcn3-ocl-bindings/libsim-ocl/build/ -lsim_ocl -lpthread -g
	./main a in.bin key.txt out.bin
