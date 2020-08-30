// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build cgo

package libc // import "modernc.org/libc"

import (
	"bufio"
	"os"
	"runtime"
	"unsafe"

	"golang.org/x/sys/unix"
	"modernc.org/libc/signal"
	"modernc.org/libc/sys/socket"
	"modernc.org/libc/sys/types"
)

/*

#cgo LDFLAGS: -lm -ldl

#include <arpa/inet.h>
#include <ctype.h>
#include <dirent.h>
#include <dlfcn.h>
#include <errno.h>
#include <fcntl.h>
#include <fts.h>
#include <grp.h>
#include <langinfo.h>
#include <locale.h>
#include <math.h>
#include <netdb.h>
#include <pwd.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <sys/resource.h>
#include <sys/select.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <sys/utsname.h>
#include <sys/wait.h>
#include <time.h>
#include <unistd.h>
#include <utime.h>

extern char **environ;

FILE *__ccgo_stdout, *__ccgo_stderr, *__ccgo_stdin;

void __ccgo_init() {
	__ccgo_stdout = stdout;
	__ccgo_stderr = stderr;
	__ccgo_stderr = stderr;
}


void __ccgo_seterrno(int err) {
	errno = err;
}

*/
import "C"

func init() {
	C.__ccgo_init()
	Xstderr = uintptr(unsafe.Pointer(C.__ccgo_stderr))
	Xstdin = uintptr(unsafe.Pointer(C.__ccgo_stdin))
	Xstdout = uintptr(unsafe.Pointer(C.__ccgo_stdout))
}

// void *realloc(void *ptr, size_t size);
func Xrealloc(t *TLS, ptr uintptr, size types.Size_t) uintptr {
	return uintptr(C.realloc(unsafe.Pointer(ptr), C.size_t(size)))
}

// void *calloc(size_t nmemb, size_t size);
func Xcalloc(t *TLS, n, size types.Size_t) uintptr {
	return uintptr(C.calloc(C.size_t(n), C.size_t(size)))
}

// void free(void *ptr);
func Xfree(t *TLS, p uintptr) { C.free(unsafe.Pointer(p)) }

// void *malloc(size_t size);
func Xmalloc(t *TLS, n types.Size_t) uintptr { return uintptr(C.malloc(C.size_t(n))) }

func Xexit(t *TLS, status int32) {
	if len(Covered) != 0 {
		buf := bufio.NewWriter(os.Stdout)
		CoverReport(buf)
		buf.Flush()
	}
	for _, v := range atExit {
		v()
	}
	C.exit(C.int(status)) // CompCert mandelbrot.c
	//TODO flush + os.Exit(int(status))
}

func Start(main func(*TLS, int32, uintptr) int32) {
	runtime.LockOSThread()
	t := NewTLS()
	argv := mustCalloc(t, types.Size_t((len(os.Args)+1)*int(uintptrSize)))
	p := argv
	for _, v := range os.Args {
		s := mustCalloc(t, types.Size_t(len(v)+1))
		copy((*RawMem)(unsafe.Pointer(s))[:len(v):len(v)], v)
		*(*uintptr)(unsafe.Pointer(p)) = s
		p += uintptrSize
	}
	Xexit(t, main(t, int32(len(os.Args)), argv))
}

func (t *TLS) setErrno(err interface{}) { //TODO -> etc.go
again:
	switch x := err.(type) {
	case int:
		C.__ccgo_seterrno(C.int(x))
	case int32:
		C.__ccgo_seterrno(C.int(x))
	case *os.PathError:
		err = x.Err
		goto again
	case unix.Errno:
		C.__ccgo_seterrno(C.int(x))
	case *os.SyscallError:
		err = x.Err
		goto again
	default:
		panic(todo("%T", x))
	}
}

// int fflush(FILE *stream);
func Xfflush(t *TLS, stream uintptr) int32 { //TODO Go version fails
	return int32(C.fflush((*C.FILE)(unsafe.Pointer(stream))))
}

// FILE *fopen64(const char *pathname, const char *mode);
func Xfopen64(t *TLS, pathname, mode uintptr) uintptr {
	return uintptr(unsafe.Pointer(C.fopen((*C.char)(unsafe.Pointer(pathname)), (*C.char)(unsafe.Pointer(mode)))))
}

// int fseek(FILE *stream, long offset, int whence);
func Xfseek(t *TLS, stream uintptr, offset long, whence int32) int32 { //TODO Go version fails
	return int32(C.fseek((*C.FILE)(unsafe.Pointer(stream)), C.long(offset), C.int(whence)))
}

// long ftell(FILE *stream);
func Xftell(t *TLS, stream uintptr) long { //TODO Go version fails
	return long(C.ftell((*C.FILE)(unsafe.Pointer(stream))))
}

// int fclose(FILE *stream);
func Xfclose(t *TLS, stream uintptr) int32 {
	return int32(C.fclose((*C.FILE)(unsafe.Pointer(stream))))
}

// size_t fread(void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfread(t *TLS, ptr uintptr, size, nmemb types.Size_t, stream uintptr) types.Size_t { //TODO Go version fails
	return types.Size_t(C.fread(unsafe.Pointer(ptr), C.size_t(size), C.size_t(nmemb), (*C.FILE)(unsafe.Pointer(stream))))
}

// int * __errno_location(void);
func X__errno_location(t *TLS) uintptr {
	return uintptr(unsafe.Pointer(C.__errno_location()))
}

// size_t fwrite(const void *ptr, size_t size, size_t nmemb, FILE *stream);
func Xfwrite(t *TLS, ptr uintptr, size, nmemb types.Size_t, stream uintptr) types.Size_t { //TODO Go version fails
	return types.Size_t(C.fwrite(unsafe.Pointer(ptr), C.size_t(size), C.size_t(nmemb), (*C.FILE)(unsafe.Pointer(stream))))
}

// int fputc(int c, FILE *stream);
func Xfputc(t *TLS, c int32, stream uintptr) int32 {
	return int32(C.fputc(C.int(c), (*C.FILE)(unsafe.Pointer(stream))))
}

// int fgetc(FILE *stream);
func Xfgetc(t *TLS, stream uintptr) int32 {
	return int32(C.fgetc((*C.FILE)(unsafe.Pointer(stream))))
}

// int fputs(const char *s, FILE *stream);
func Xfputs(t *TLS, s, stream uintptr) int32 {
	return int32(C.fputs((*C.char)(unsafe.Pointer(s)), (*C.FILE)(unsafe.Pointer(stream))))
}

// struct servent *getservbyname(const char *name, const char *proto);
func Xgetservbyname(t *TLS, name, proto uintptr) uintptr {
	return uintptr(unsafe.Pointer(C.getservbyname((*C.char)(unsafe.Pointer(name)), (*C.char)(unsafe.Pointer(proto)))))
}

// int getaddrinfo(const char *node, const char *service, const struct addrinfo *hints, struct addrinfo **res);
func Xgetaddrinfo(t *TLS, node, service, hints, res uintptr) int32 { //TODO not needed by sqlite
	return int32(C.getaddrinfo((*C.char)(unsafe.Pointer(node)), (*C.char)(unsafe.Pointer(service)), (*C.struct_addrinfo)(unsafe.Pointer(hints)), (**C.struct_addrinfo)(unsafe.Pointer(res))))
}

// FILE *fdopen(int fd, const char *mode);
func Xfdopen(t *TLS, fd int32, mode uintptr) uintptr {
	return uintptr(unsafe.Pointer(C.fdopen(C.int(fd), (*C.char)(unsafe.Pointer(mode)))))
}

// void _exit(int status);
func X_exit(t *TLS, status int32) {
	C._exit(C.int(status))
}

// void freeaddrinfo(struct addrinfo *res);
func Xfreeaddrinfo(t *TLS, res uintptr) {
	C.freeaddrinfo((*C.struct_addrinfo)(unsafe.Pointer(res)))
}

// int getnameinfo(const struct sockaddr *addr, socklen_t addrlen, char *host, socklen_t hostlen, char *serv, socklen_t servlen, int flags);
func Xgetnameinfo(t *TLS, addr uintptr, addrlen socket.Socklen_t, host uintptr, hostlen socket.Socklen_t, serv uintptr, servlen socket.Socklen_t, flags int32) int32 {
	return int32(C.getnameinfo(
		(*C.struct_sockaddr)(unsafe.Pointer(addr)),
		C.socklen_t(addrlen),
		(*C.char)(unsafe.Pointer(host)),
		C.socklen_t(hostlen),
		(*C.char)(unsafe.Pointer(serv)),
		C.socklen_t(servlen),
		C.int(flags),
	))
}

// struct hostent *gethostbyname(const char *name);
func Xgethostbyname(t *TLS, name uintptr) uintptr {
	return uintptr(unsafe.Pointer(C.gethostbyname((*C.char)(unsafe.Pointer(name)))))
}

// struct hostent *gethostbyaddr(const void *addr, socklen_t len, int type);
func Xgethostbyaddr(t *TLS, addr uintptr, len socket.Socklen_t, type1 int32) uintptr {
	return uintptr(unsafe.Pointer(C.gethostbyaddr(unsafe.Pointer(addr), C.socklen_t(len), C.int(type1))))
}

// int ferror(FILE *stream);
func Xferror(t *TLS, stream uintptr) int32 {
	return int32(C.ferror((*C.FILE)(unsafe.Pointer(stream))))
}

func Environ() uintptr {
	return uintptr(unsafe.Pointer(C.environ))
}

func EnvironP() uintptr {
	return uintptr(unsafe.Pointer(&C.environ))
}

var abortSigaction signal.Sigaction

func Xabort(t *TLS) {
	abortSigaction = signal.Sigaction{
		F__sigaction_handler: struct{ Fsa_handler signal.X__sighandler_t }{Fsa_handler: signal.SIG_DFL},
	}
	if C.sigaction(signal.SIGABRT, (*C.struct_sigaction)(unsafe.Pointer(&abortSigaction)), nil) != 0 {
		panic(todo(""))
	}

	C.abort()
}
