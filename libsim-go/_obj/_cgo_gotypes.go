// Created by cgo - DO NOT EDIT

package gcn3simocl

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_int int32

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


//go:linkname _cgo_runtime_gobytes runtime.gobytes
func _cgo_runtime_gobytes(unsafe.Pointer, int) []byte

func _Cfunc_GoBytes(p unsafe.Pointer, l _Ctype_int) []byte {
	return _cgo_runtime_gobytes(p, int(l))
}
//go:cgo_export_dynamic initializeSimulator
//go:linkname _cgoexp_281341e1c880_initializeSimulator _cgoexp_281341e1c880_initializeSimulator
//go:cgo_export_static _cgoexp_281341e1c880_initializeSimulator
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_initializeSimulator(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_initializeSimulator
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_initializeSimulator() {
	initializeSimulator()
}
//go:cgo_export_dynamic gcn3GetPlatformIDs
//go:linkname _cgoexp_281341e1c880_gcn3GetPlatformIDs _cgoexp_281341e1c880_gcn3GetPlatformIDs
//go:cgo_export_static _cgoexp_281341e1c880_gcn3GetPlatformIDs
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3GetPlatformIDs(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3GetPlatformIDs
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3GetPlatformIDs() (r0 int) {
	return gcn3GetPlatformIDs()
}
//go:cgo_export_dynamic gcn3GetDeviceIDs
//go:linkname _cgoexp_281341e1c880_gcn3GetDeviceIDs _cgoexp_281341e1c880_gcn3GetDeviceIDs
//go:cgo_export_static _cgoexp_281341e1c880_gcn3GetDeviceIDs
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3GetDeviceIDs(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3GetDeviceIDs
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3GetDeviceIDs() (r0 int) {
	return gcn3GetDeviceIDs()
}
//go:cgo_export_dynamic gcn3CreateContext
//go:linkname _cgoexp_281341e1c880_gcn3CreateContext _cgoexp_281341e1c880_gcn3CreateContext
//go:cgo_export_static _cgoexp_281341e1c880_gcn3CreateContext
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3CreateContext(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3CreateContext
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3CreateContext() (r0 int) {
	return gcn3CreateContext()
}
//go:cgo_export_dynamic gcn3CreateProgramWithSource
//go:linkname _cgoexp_281341e1c880_gcn3CreateProgramWithSource _cgoexp_281341e1c880_gcn3CreateProgramWithSource
//go:cgo_export_static _cgoexp_281341e1c880_gcn3CreateProgramWithSource
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3CreateProgramWithSource(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3CreateProgramWithSource
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3CreateProgramWithSource(p0 int, p1 string) (r0 int) {
	return gcn3CreateProgramWithSource(p0, p1)
}
//go:cgo_export_dynamic gcn3BuildProgram
//go:linkname _cgoexp_281341e1c880_gcn3BuildProgram _cgoexp_281341e1c880_gcn3BuildProgram
//go:cgo_export_static _cgoexp_281341e1c880_gcn3BuildProgram
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3BuildProgram(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3BuildProgram
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3BuildProgram(p0 int) (r0 int) {
	return gcn3BuildProgram(p0)
}
//go:cgo_export_dynamic gcn3CreateKernel
//go:linkname _cgoexp_281341e1c880_gcn3CreateKernel _cgoexp_281341e1c880_gcn3CreateKernel
//go:cgo_export_static _cgoexp_281341e1c880_gcn3CreateKernel
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3CreateKernel(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3CreateKernel
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3CreateKernel(p0 int, p1 string) (r0 int) {
	return gcn3CreateKernel(p0, p1)
}
//go:cgo_export_dynamic gcn3CreateBuffer
//go:linkname _cgoexp_281341e1c880_gcn3CreateBuffer _cgoexp_281341e1c880_gcn3CreateBuffer
//go:cgo_export_static _cgoexp_281341e1c880_gcn3CreateBuffer
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3CreateBuffer(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3CreateBuffer
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3CreateBuffer(p0 int, p1 int) (r0 int) {
	return gcn3CreateBuffer(p0, p1)
}
//go:cgo_export_dynamic gcn3EnqueueWriteBuffer
//go:linkname _cgoexp_281341e1c880_gcn3EnqueueWriteBuffer _cgoexp_281341e1c880_gcn3EnqueueWriteBuffer
//go:cgo_export_static _cgoexp_281341e1c880_gcn3EnqueueWriteBuffer
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3EnqueueWriteBuffer(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3EnqueueWriteBuffer
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3EnqueueWriteBuffer(p0 int, p1 int, p2 unsafe.Pointer) (r0 int) {
	return gcn3EnqueueWriteBuffer(p0, p1, p2)
}
//go:cgo_export_dynamic gcn3EnqueueReadBuffer
//go:linkname _cgoexp_281341e1c880_gcn3EnqueueReadBuffer _cgoexp_281341e1c880_gcn3EnqueueReadBuffer
//go:cgo_export_static _cgoexp_281341e1c880_gcn3EnqueueReadBuffer
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3EnqueueReadBuffer(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3EnqueueReadBuffer
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3EnqueueReadBuffer(p0 int, p1 int, p2 unsafe.Pointer) (r0 int) {
	return gcn3EnqueueReadBuffer(p0, p1, p2)
}
//go:cgo_export_dynamic gcn3SetKernelArg
//go:linkname _cgoexp_281341e1c880_gcn3SetKernelArg _cgoexp_281341e1c880_gcn3SetKernelArg
//go:cgo_export_static _cgoexp_281341e1c880_gcn3SetKernelArg
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3SetKernelArg(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3SetKernelArg
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3SetKernelArg(p0 int, p1 int, p2 int, p3 unsafe.Pointer) (r0 int) {
	return gcn3SetKernelArg(p0, p1, p2, p3)
}
//go:cgo_export_dynamic gcn3LaunchKernel
//go:linkname _cgoexp_281341e1c880_gcn3LaunchKernel _cgoexp_281341e1c880_gcn3LaunchKernel
//go:cgo_export_static _cgoexp_281341e1c880_gcn3LaunchKernel
//go:nosplit
//go:norace
func _cgoexp_281341e1c880_gcn3LaunchKernel(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_281341e1c880_gcn3LaunchKernel
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_281341e1c880_gcn3LaunchKernel(p0 int, p1 unsafe.Pointer, p2 unsafe.Pointer) (r0 int) {
	return gcn3LaunchKernel(p0, p1, p2)
}
