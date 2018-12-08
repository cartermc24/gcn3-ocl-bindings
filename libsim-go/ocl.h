/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "/home/carter/simulator/gcn3-ocl-bindings/libsim-go/ocl.go"

 
#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef _GoString_ GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


//
// Helpers
//

extern void initializeSimulator();

//
// OpenCL API
//

extern GoInt gcn3GetPlatformIDs();

extern GoInt gcn3GetDeviceIDs();

extern GoInt gcn3GetKernelInfo(GoInt p0, GoInt p1, GoUint64 p2, void* p3, void* p4);

extern GoInt gcn3GetProgramBuildInfo(GoInt p0, GoInt p1, GoInt p2, GoUint64 p3, void* p4, void* p5);

extern GoInt gcn3GetDeviceInfo(GoInt p0, GoInt p1, GoUint64 p2, void* p3, void* p4);

extern GoInt gcn3CreateContext();

// Returns Kernel ID

extern GoInt gcn3CreateProgramWithSource(GoInt p0, GoString p1);

extern GoInt gcn3BuildProgram(GoInt p0);

extern GoInt gcn3CreateKernelsInProgram(GoInt p0, GoUint p1, void* p2, void* p3, GoInt p4);

extern GoInt gcn3CreateKernel(GoInt p0, GoString p1);

// Returns Buffer ID
// Size in bytes

extern GoInt gcn3CreateBuffer(GoInt p0, GoInt p1);

extern GoInt gcn3EnqueueWriteBuffer(GoInt p0, GoInt p1, void* p2);

extern GoInt gcn3EnqueueReadBuffer(GoInt p0, GoInt p1, void* p2);

extern GoInt gcn3SetKernelArg(GoInt p0, GoInt p1, GoInt p2, void* p3);

// global_work_size is type uint32
// local_work_size is type uint16

extern GoInt gcn3LaunchKernel(GoInt p0, void* p1, void* p2);

extern GoInt gcn3ReleaseProgram(GoInt p0);

#ifdef __cplusplus
}
#endif
