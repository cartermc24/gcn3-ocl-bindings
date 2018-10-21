/* Created by cgo - DO NOT EDIT. */
#include <stdlib.h>
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int, __SIZE_TYPE__), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done();
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

extern void _cgoexp_281341e1c880_initializeSimulator(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void initializeSimulator()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char unused;
	} __attribute__((__packed__, __gcc_struct__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_initializeSimulator, &a, 0, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}
extern void _cgoexp_281341e1c880_gcn3GetPlatformIDs(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3GetPlatformIDs()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3GetPlatformIDs, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3GetDeviceIDs(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3GetDeviceIDs()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3GetDeviceIDs, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3CreateContext(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3CreateContext()
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3CreateContext, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3CreateProgramWithSource(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3CreateProgramWithSource(GoInt p0, GoString p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoString p1;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3CreateProgramWithSource, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3BuildProgram(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3BuildProgram(GoInt p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3BuildProgram, &a, 16, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3CreateKernel(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3CreateKernel(GoInt p0, GoString p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoString p1;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3CreateKernel, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3CreateBuffer(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3CreateBuffer(GoInt p0, GoInt p1)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoInt p1;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3CreateBuffer, &a, 24, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3EnqueueWriteBuffer(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3EnqueueWriteBuffer(GoInt p0, GoInt p1, void* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoInt p1;
		void* p2;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3EnqueueWriteBuffer, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3EnqueueReadBuffer(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3EnqueueReadBuffer(GoInt p0, GoInt p1, void* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoInt p1;
		void* p2;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3EnqueueReadBuffer, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3SetKernelArg(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3SetKernelArg(GoInt p0, GoInt p1, GoInt p2, void* p3)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		GoInt p1;
		GoInt p2;
		void* p3;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	a.p3 = p3;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3SetKernelArg, &a, 40, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
extern void _cgoexp_281341e1c880_gcn3LaunchKernel(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
GoInt gcn3LaunchKernel(GoInt p0, void* p1, void* p2)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		GoInt p0;
		void* p1;
		void* p2;
		GoInt r0;
	} __attribute__((__packed__, __gcc_struct__)) a;
	a.p0 = p0;
	a.p1 = p1;
	a.p2 = p2;
	_cgo_tsan_release();
	crosscall2(_cgoexp_281341e1c880_gcn3LaunchKernel, &a, 32, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
	return a.r0;
}
