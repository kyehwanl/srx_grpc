/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "/opt/project/gobgp_test/gowork/src/srx_grpc/server/srx_grpc_server.go"






//#cgo LDFLAGS: -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -lgrpc_service -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/server -L/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -lSRxProxy -Wl,-rpath -Wl,/opt/project/gobgp_test/gowork/src/srx_grpc/srx/test_install/lib64/srx -Wl,--unresolved-symbols=ignore-all

#include <stdio.h>
#include "srx/srx_api.h"
#include "server/grpc_service.h"



extern void cb_proxy(int f, void* user_data);

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


extern void cb_proxy(int p0, void* p1);

extern void cb_proxyGoodBye(SRXPROXY_GOODBYE p0);

extern void cb_proxyStream(int p0, void* p1);

extern void Serve();

#ifdef __cplusplus
}
#endif
