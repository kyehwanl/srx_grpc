/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */


#line 1 "cgo-builtin-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

typedef struct { const char *p; ptrdiff_t n; } _GoString_;

#endif

/* Start of preamble from import "C" comments.  */


#line 3 "/opt/project/gobgp_test/gowork/src/srx_grpc/client/srx_grpc_client.go"




#include <stdlib.h>
#include "shared/srx_packets.h"
#include "client/grpc_client_service.h"

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


extern GoUint8 InitWorkerPool();

extern GoUint8 InitSRxGrpc(GoString p0);

extern GoUint32 Run(GoSlice p0);

/* Return type for RunProxyHello */
struct RunProxyHello_return {
	unsigned char* r0;
	GoUint32 r1;
};

extern struct RunProxyHello_return RunProxyHello(GoSlice p0);

extern GoUint8 RunProxyGoodBye(SRXPROXY_GOODBYE p0, GoUint32 p1);

extern GoUint32 RunProxyGoodByeStream(GoSlice p0, GoUint32 p1);

extern GoUint32 RunProxyStream(GoSlice p0, GoUint32 p1);

extern GoUint32 RunStream(GoSlice p0);

extern GoUint32 RunProxyVerify(GoSlice p0, GoUint32 p1);

#ifdef __cplusplus
}
#endif