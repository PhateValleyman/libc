// Code generated by 'ccgo errno/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o errno/errno_windows_amd64.go -pkgname errno', DO NOT EDIT.

package errno

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	DUMMYSTRUCTNAME                                 = 0
	DUMMYSTRUCTNAME1                                = 0
	DUMMYSTRUCTNAME2                                = 0
	DUMMYSTRUCTNAME3                                = 0
	DUMMYSTRUCTNAME4                                = 0
	DUMMYSTRUCTNAME5                                = 0
	DUMMYUNIONNAME                                  = 0
	DUMMYUNIONNAME1                                 = 0
	DUMMYUNIONNAME2                                 = 0
	DUMMYUNIONNAME3                                 = 0
	DUMMYUNIONNAME4                                 = 0
	DUMMYUNIONNAME5                                 = 0
	DUMMYUNIONNAME6                                 = 0
	DUMMYUNIONNAME7                                 = 0
	DUMMYUNIONNAME8                                 = 0
	DUMMYUNIONNAME9                                 = 0
	E2BIG                                           = 7
	EACCES                                          = 13
	EADDRINUSE                                      = 100
	EADDRNOTAVAIL                                   = 101
	EAFNOSUPPORT                                    = 102
	EAGAIN                                          = 11
	EALREADY                                        = 103
	EBADF                                           = 9
	EBADMSG                                         = 104
	EBUSY                                           = 16
	ECANCELED                                       = 105
	ECHILD                                          = 10
	ECONNABORTED                                    = 106
	ECONNREFUSED                                    = 107
	ECONNRESET                                      = 108
	EDEADLK                                         = 36
	EDEADLOCK                                       = 36
	EDESTADDRREQ                                    = 109
	EDOM                                            = 33
	EEXIST                                          = 17
	EFAULT                                          = 14
	EFBIG                                           = 27
	EHOSTUNREACH                                    = 110
	EIDRM                                           = 111
	EILSEQ                                          = 42
	EINPROGRESS                                     = 112
	EINTR                                           = 4
	EINVAL                                          = 22
	EIO                                             = 5
	EISCONN                                         = 113
	EISDIR                                          = 21
	ELOOP                                           = 114
	EMFILE                                          = 24
	EMLINK                                          = 31
	EMSGSIZE                                        = 115
	ENAMETOOLONG                                    = 38
	ENETDOWN                                        = 116
	ENETRESET                                       = 117
	ENETUNREACH                                     = 118
	ENFILE                                          = 23
	ENOBUFS                                         = 119
	ENODATA                                         = 120
	ENODEV                                          = 19
	ENOENT                                          = 2
	ENOEXEC                                         = 8
	ENOFILE                                         = 2
	ENOLCK                                          = 39
	ENOLINK                                         = 121
	ENOMEM                                          = 12
	ENOMSG                                          = 122
	ENOPROTOOPT                                     = 123
	ENOSPC                                          = 28
	ENOSR                                           = 124
	ENOSTR                                          = 125
	ENOSYS                                          = 40
	ENOTCONN                                        = 126
	ENOTDIR                                         = 20
	ENOTEMPTY                                       = 41
	ENOTRECOVERABLE                                 = 127
	ENOTSOCK                                        = 128
	ENOTSUP                                         = 129
	ENOTTY                                          = 25
	ENXIO                                           = 6
	EOPNOTSUPP                                      = 130
	EOVERFLOW                                       = 132
	EOWNERDEAD                                      = 133
	EPERM                                           = 1
	EPIPE                                           = 32
	EPROTO                                          = 134
	EPROTONOSUPPORT                                 = 135
	EPROTOTYPE                                      = 136
	ERANGE                                          = 34
	EROFS                                           = 30
	ESPIPE                                          = 29
	ESRCH                                           = 3
	ETIME                                           = 137
	ETIMEDOUT                                       = 138
	ETXTBSY                                         = 139
	EWOULDBLOCK                                     = 140
	EXDEV                                           = 18
	MINGW_DDK_H                                     = 0
	MINGW_DDRAW_VERSION                             = 7
	MINGW_HAS_DDK_H                                 = 1
	MINGW_HAS_DDRAW_H                               = 1
	MINGW_HAS_SECURE_API                            = 1
	MINGW_SDK_INIT                                  = 0
	STRUNCATE                                       = 80
	UNALIGNED                                       = 0
	USE___UUIDOF                                    = 0
	WIN32                                           = 1
	WIN64                                           = 1
	WINNT                                           = 1
	X_AGLOBAL                                       = 0
	X_ANONYMOUS_STRUCT                              = 0
	X_ANONYMOUS_UNION                               = 0
	X_ARGMAX                                        = 100
	X_CONST_RETURN                                  = 0
	X_CRTIMP                                        = 0
	X_CRTIMP2                                       = 0
	X_CRTIMP_ALTERNATIVE                            = 0
	X_CRTIMP_NOIA64                                 = 0
	X_CRTIMP_PURE                                   = 0
	X_CRTNOALIAS                                    = 0
	X_CRTRESTRICT                                   = 0
	X_CRT_ALTERNATIVE_IMPORTED                      = 0
	X_CRT_ERRNO_DEFINED                             = 0
	X_CRT_MANAGED_HEAP_DEPRECATE                    = 0
	X_CRT_PACKING                                   = 8
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES          = 0
	X_CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES_MEMORY   = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES        = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT  = 0
	X_CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_MEMORY = 0
	X_DLL                                           = 0
	X_ERRCODE_DEFINED                               = 0
	X_FILE_OFFSET_BITS                              = 64
	X_INC_CRTDEFS                                   = 0
	X_INC_CRTDEFS_MACRO                             = 0
	X_INC_ERRNO                                     = 0
	X_INC_MINGW_SECAPI                              = 0
	X_INC_VADEFS                                    = 0
	X_INC__MINGW_H                                  = 0
	X_INT128_DEFINED                                = 0
	X_INTEGRAL_MAX_BITS                             = 64
	X_INTPTR_T_DEFINED                              = 0
	X_MCRTIMP                                       = 0
	X_MRTIMP2                                       = 0
	X_MT                                            = 0
	X_M_AMD64                                       = 100
	X_M_X64                                         = 100
	X_PGLOBAL                                       = 0
	X_PTRDIFF_T_                                    = 0
	X_PTRDIFF_T_DEFINED                             = 0
	X_RSIZE_T_DEFINED                               = 0
	X_SECURECRT_ERRCODE_VALUES_DEFINED              = 0
	X_SECURECRT_FILL_BUFFER_PATTERN                 = 0xFD
	X_SIZE_T_DEFINED                                = 0
	X_SSIZE_T_DEFINED                               = 0
	X_TAGLC_ID_DEFINED                              = 0
	X_THREADLOCALEINFO                              = 0
	X_TIME32_T_DEFINED                              = 0
	X_TIME64_T_DEFINED                              = 0
	X_TIME_T_DEFINED                                = 0
	X_UINTPTR_T_DEFINED                             = 0
	X_VA_LIST_DEFINED                               = 0
	X_W64                                           = 0
	X_WCHAR_T_DEFINED                               = 0
	X_WCTYPE_T_DEFINED                              = 0
	X_WIN32                                         = 1
	X_WIN32_WINNT                                   = 0x502
	X_WIN64                                         = 1
	X_WINT_T                                        = 0
	X_cdecl                                         = 0
	X_fastcall                                      = 0
	X_stdcall                                       = 0
	X_thiscall                                      = 0
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = uint16 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:47:14 */
type X__float128 = float64        /* <builtin>:48:21 */

type Va_list = X__builtin_va_list /* <builtin>:51:27 */
// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// This macro holds an monotonic increasing value, which indicates
//    a specific fix/patch is present on trunk.  This value isn't related to
//    minor/major version-macros.  It is increased on demand, if a big
//    fix was applied to trunk.  This macro gets just increased on trunk.  For
//    other branches its value won't be modified.

// mingw.org's version macros: these make gcc to define
//    MINGW32_SUPPORTS_MT_EH and to use the _CRT_MT global
//    and the __mingwthr_key_dtor() function from the MinGW
//    CRT in its private gthr-win32.h header.

// Set VC specific compiler target macros.

// MS does not prefix symbols by underscores for 64-bit.
// As we have to support older gcc version, which are using underscores
//       as symbol prefix for x64, we have to check here for the user label
//       prefix defined by gcc.

// Special case nameless struct/union.

// MinGW-w64 has some additional C99 printf/scanf feature support.
//    So we add some helper macros to ease recognition of them.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// http://msdn.microsoft.com/en-us/library/ms175759%28v=VS.100%29.aspx
// Templates won't work in C, will break if secure API is not enabled, disabled

// https://blogs.msdn.com/b/sdl/archive/2010/02/16/vc-2010-and-memcpy.aspx?Redirected=true
// fallback on default implementation if we can't know the size of the destination

// Include _cygwin.h if we're building a Cygwin application.

// Target specific macro replacement for type "long".  In the Windows API,
//    the type long is always 32 bit, even if the target is 64 bit (LLP64).
//    On 64 bit Cygwin, the type long is 64 bit (LP64).  So, to get the right
//    sized definitions and declarations, all usage of type long in the Windows
//    headers have to be replaced by the below defined macro __LONG32.

// C/C++ specific language defines.

// Note the extern. This is needed to work around GCC's
// limitations in handling dllimport attribute.

// Attribute `nonnull' was valid as of gcc 3.3.  We don't use GCC's
//    variadiac macro facility, because variadic macros cause syntax
//    errors with  --traditional-cpp.

//  High byte is the major version, low byte is the minor.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

// *
// This file has no copyright assigned and is placed in the Public Domain.
// This file is part of the mingw-w64 runtime package.
// No warranty is given; refer to the file DISCLAIMER.PD within this package.

type X__gnuc_va_list = X__builtin_va_list /* vadefs.h:24:29 */

type Ssize_t = int64 /* crtdefs.h:45:35 */

type Rsize_t = Size_t /* crtdefs.h:52:16 */

type Intptr_t = int64 /* crtdefs.h:62:35 */

type Uintptr_t = uint64 /* crtdefs.h:75:44 */

type Wint_t = uint16   /* crtdefs.h:106:24 */
type Wctype_t = uint16 /* crtdefs.h:107:24 */

type Errno_t = int32 /* crtdefs.h:113:13 */

type X__time32_t = int32 /* crtdefs.h:118:14 */

type X__time64_t = int64 /* crtdefs.h:123:35 */

type Time_t = X__time64_t /* crtdefs.h:138:20 */

type Threadlocaleinfostruct = struct {
	Frefcount      int32
	Flc_codepage   uint32
	Flc_collate_cp uint32
	Flc_handle     [6]uint32
	Flc_id         [6]LC_ID
	Flc_category   [6]struct {
		Flocale    uintptr
		Fwlocale   uintptr
		Frefcount  uintptr
		Fwrefcount uintptr
	}
	Flc_clike            int32
	Fmb_cur_max          int32
	Flconv_intl_refcount uintptr
	Flconv_num_refcount  uintptr
	Flconv_mon_refcount  uintptr
	Flconv               uintptr
	Fctype1_refcount     uintptr
	Fctype1              uintptr
	Fpctype              uintptr
	Fpclmap              uintptr
	Fpcumap              uintptr
	Flc_time_curr        uintptr
} /* crtdefs.h:422:1 */

type Pthreadlocinfo = uintptr /* crtdefs.h:424:39 */
type Pthreadmbcinfo = uintptr /* crtdefs.h:425:36 */

type Localeinfo_struct = struct {
	Flocinfo Pthreadlocinfo
	Fmbcinfo Pthreadmbcinfo
} /* crtdefs.h:428:9 */

type X_locale_tstruct = Localeinfo_struct /* crtdefs.h:431:3 */
type X_locale_t = uintptr                 /* crtdefs.h:431:19 */

type TagLC_ID = struct {
	FwLanguage uint16
	FwCountry  uint16
	FwCodePage uint16
} /* crtdefs.h:422:1 */

type LC_ID = TagLC_ID  /* crtdefs.h:439:3 */
type LPLC_ID = uintptr /* crtdefs.h:439:9 */

type Threadlocinfo = Threadlocaleinfostruct /* crtdefs.h:468:3 */

// Posix thread extensions.

// Extension defined as by report VC 10+ defines error-numbers.

// Defined as WSAETIMEDOUT.

var _ int8 /* gen.c:2:13: */
