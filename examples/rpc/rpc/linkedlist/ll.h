/*
 * Please do not edit this file.
 * It was generated using rpcgen.
 */

#ifndef _LL_H_RPCGEN
#define _LL_H_RPCGEN

#include <rpc/rpc.h>


#ifdef __cplusplus
extern "C" {
#endif


struct foo {
	int x;
	struct foo *next;
};
typedef struct foo foo;

#define LL_PROG 555553555
#define LL_VERSION 1

#if defined(__STDC__) || defined(__cplusplus)
#define SUM 1
extern  int * sum_1(foo *, CLIENT *);
extern  int * sum_1_svc(foo *, struct svc_req *);
extern int ll_prog_1_freeresult (SVCXPRT *, xdrproc_t, caddr_t);

#else /* K&R C */
#define SUM 1
extern  int * sum_1();
extern  int * sum_1_svc();
extern int ll_prog_1_freeresult ();
#endif /* K&R C */

/* the xdr functions */

#if defined(__STDC__) || defined(__cplusplus)
extern  bool_t xdr_foo (XDR *, foo*);

#else /* K&R C */
extern bool_t xdr_foo ();

#endif /* K&R C */

#ifdef __cplusplus
}
#endif

#endif /* !_LL_H_RPCGEN */