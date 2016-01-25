/*
 * Please do not edit this file.
 * It was generated using rpcgen.
 */

#include <memory.h> /* for memset */
#include "userlookup.h"
#define NOTFOUND 0
#define FOUND 1

/* Default timeout can be changed using clnt_control() */
static struct timeval TIMEOUT = { 25, 0 };

uid_retval *
byname_1(username *argp, CLIENT *clnt)
{
	static uid_retval clnt_res;

	memset((char *)&clnt_res, 0, sizeof(clnt_res));
	if (clnt_call (clnt, byname,
		(xdrproc_t) xdr_username, (caddr_t) argp,
		(xdrproc_t) xdr_uid_retval, (caddr_t) &clnt_res,
		TIMEOUT) != RPC_SUCCESS) {
		return (NULL);
	}
	return (&clnt_res);
}

uname_retval *
bynum_1(int *argp, CLIENT *clnt)
{
	static uname_retval clnt_res;

	memset((char *)&clnt_res, 0, sizeof(clnt_res));
	if (clnt_call (clnt, bynum,
		(xdrproc_t) xdr_int, (caddr_t) argp,
		(xdrproc_t) xdr_uname_retval, (caddr_t) &clnt_res,
		TIMEOUT) != RPC_SUCCESS) {
		return (NULL);
	}
	return (&clnt_res);
}