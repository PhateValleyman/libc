// Code generated by 'ccgo limits/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o limits/limits_linux_amd64.go -pkgname limits', DO NOT EDIT.

package limits

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ unsafe.Pointer

const (
	AIO_PRIO_DELTA_MAX                   = 20
	BC_BASE_MAX                          = 99
	BC_DIM_MAX                           = 2048
	BC_SCALE_MAX                         = 99
	BC_STRING_MAX                        = 1000
	CHARCLASS_NAME_MAX                   = 2048
	CHAR_BIT                             = 8
	CHAR_MAX                             = 127
	CHAR_MIN                             = -128
	COLL_WEIGHTS_MAX                     = 255
	DELAYTIMER_MAX                       = 2147483647
	EXPR_NEST_MAX                        = 32
	HOST_NAME_MAX                        = 64
	INT_MAX                              = 2147483647
	INT_MIN                              = -2147483648
	LINE_MAX                             = 2048
	LLONG_MAX                            = 9223372036854775807
	LLONG_MIN                            = -9223372036854775808
	LOGIN_NAME_MAX                       = 256
	LONG_MAX                             = 9223372036854775807
	LONG_MIN                             = -9223372036854775808
	MAX_CANON                            = 255
	MAX_INPUT                            = 255
	MB_LEN_MAX                           = 16
	MQ_PRIO_MAX                          = 32768
	NAME_MAX                             = 255
	NGROUPS_MAX                          = 65536
	PATH_MAX                             = 4096
	PIPE_BUF                             = 4096
	PTHREAD_DESTRUCTOR_ITERATIONS        = 4
	PTHREAD_KEYS_MAX                     = 1024
	PTHREAD_STACK_MIN                    = 16384
	RE_DUP_MAX                           = 32767
	RTSIG_MAX                            = 32
	SCHAR_MAX                            = 127
	SCHAR_MIN                            = -128
	SEM_VALUE_MAX                        = 2147483647
	SHRT_MAX                             = 32767
	SHRT_MIN                             = -32768
	SSIZE_MAX                            = 9223372036854775807
	TTY_NAME_MAX                         = 32
	UCHAR_MAX                            = 255
	UINT_MAX                             = 4294967295
	ULLONG_MAX                           = 18446744073709551615
	ULONG_MAX                            = 18446744073709551615
	USHRT_MAX                            = 65535
	XATTR_LIST_MAX                       = 65536
	XATTR_NAME_MAX                       = 255
	XATTR_SIZE_MAX                       = 65536
	X_ATFILE_SOURCE                      = 1
	X_BITS_POSIX1_LIM_H                  = 1
	X_BITS_POSIX2_LIM_H                  = 1
	X_DEFAULT_SOURCE                     = 1
	X_FEATURES_H                         = 1
	X_FILE_OFFSET_BITS                   = 64
	X_GCC_LIMITS_H_                      = 0
	X_LIBC_LIMITS_H_                     = 1
	X_LIMITS_H___                        = 0
	X_LINUX_LIMITS_H                     = 0
	X_LP64                               = 1
	X_POSIX2_BC_BASE_MAX                 = 99
	X_POSIX2_BC_DIM_MAX                  = 2048
	X_POSIX2_BC_SCALE_MAX                = 99
	X_POSIX2_BC_STRING_MAX               = 1000
	X_POSIX2_CHARCLASS_NAME_MAX          = 14
	X_POSIX2_COLL_WEIGHTS_MAX            = 2
	X_POSIX2_EXPR_NEST_MAX               = 32
	X_POSIX2_LINE_MAX                    = 2048
	X_POSIX2_RE_DUP_MAX                  = 255
	X_POSIX_AIO_LISTIO_MAX               = 2
	X_POSIX_AIO_MAX                      = 1
	X_POSIX_ARG_MAX                      = 4096
	X_POSIX_CHILD_MAX                    = 25
	X_POSIX_CLOCKRES_MIN                 = 20000000
	X_POSIX_C_SOURCE                     = 200809
	X_POSIX_DELAYTIMER_MAX               = 32
	X_POSIX_HOST_NAME_MAX                = 255
	X_POSIX_LINK_MAX                     = 8
	X_POSIX_LOGIN_NAME_MAX               = 9
	X_POSIX_MAX_CANON                    = 255
	X_POSIX_MAX_INPUT                    = 255
	X_POSIX_MQ_OPEN_MAX                  = 8
	X_POSIX_MQ_PRIO_MAX                  = 32
	X_POSIX_NAME_MAX                     = 14
	X_POSIX_NGROUPS_MAX                  = 8
	X_POSIX_OPEN_MAX                     = 20
	X_POSIX_PATH_MAX                     = 256
	X_POSIX_PIPE_BUF                     = 512
	X_POSIX_RE_DUP_MAX                   = 255
	X_POSIX_RTSIG_MAX                    = 8
	X_POSIX_SEM_NSEMS_MAX                = 256
	X_POSIX_SEM_VALUE_MAX                = 32767
	X_POSIX_SIGQUEUE_MAX                 = 32
	X_POSIX_SOURCE                       = 1
	X_POSIX_SSIZE_MAX                    = 32767
	X_POSIX_STREAM_MAX                   = 8
	X_POSIX_SYMLINK_MAX                  = 255
	X_POSIX_SYMLOOP_MAX                  = 8
	X_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
	X_POSIX_THREAD_KEYS_MAX              = 128
	X_POSIX_THREAD_THREADS_MAX           = 64
	X_POSIX_TIMER_MAX                    = 32
	X_POSIX_TTY_NAME_MAX                 = 9
	X_POSIX_TZNAME_MAX                   = 6
	X_STDC_PREDEF_H                      = 1
	X_SYS_CDEFS_H                        = 1
	Linux                                = 1
	Unix                                 = 1
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = [2]int64   /* <builtin>:21:24 */ //TODO
type X__uint128_t = [2]uint64 /* <builtin>:22:25 */ //TODO

type X__builtin_va_list = uintptr /* <builtin>:42:14 */
type X__float128 = float64        /* <builtin>:43:21 */
// Copyright (C) 1992-2018 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation; either version 3, or (at your option) any later
// version.
//
// GCC is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
// for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// This administrivia gets added to the beginning of limits.h
//    if the system has its own version of limits.h.

// We use _GCC_LIMITS_H_ because we want this not to match
//    any macros that the system's limits.h uses for its own purposes.

// Use "..." so that we find syslimits.h only in this same directory.
// syslimits.h stands for the system's own limits.h file.
//    If we can use it ok unmodified, then we install this text.
//    If fixincludes fixes it, then the fixed version is installed
//    instead of this text.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

//	ISO C99 Standard: 7.10/5.2.4.2.1 Sizes of integer types	<limits.h>

// Handle feature test macros at the start of a header.
//    Copyright (C) 2016-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// This header is internal to glibc and should not be included outside
//    of glibc headers.  Headers including it must define
//    __GLIBC_INTERNAL_STARTING_HEADER_IMPLEMENTATION first.  This header
//    cannot have multiple include guards because ISO C feature test
//    macros depend on the definition of the macro when an affected
//    header is included, not when the first system header is
//    included.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// These are defined by the user (or the compiler)
//    to specify the desired environment:
//
//    __STRICT_ANSI__	ISO Standard C.
//    _ISOC99_SOURCE	Extensions to ISO C89 from ISO C99.
//    _ISOC11_SOURCE	Extensions to ISO C99 from ISO C11.
//    __STDC_WANT_LIB_EXT2__
// 			Extensions to ISO C99 from TR 27431-2:2010.
//    __STDC_WANT_IEC_60559_BFP_EXT__
// 			Extensions to ISO C11 from TS 18661-1:2014.
//    __STDC_WANT_IEC_60559_FUNCS_EXT__
// 			Extensions to ISO C11 from TS 18661-4:2015.
//    __STDC_WANT_IEC_60559_TYPES_EXT__
// 			Extensions to ISO C11 from TS 18661-3:2015.
//
//    _POSIX_SOURCE	IEEE Std 1003.1.
//    _POSIX_C_SOURCE	If ==1, like _POSIX_SOURCE; if >=2 add IEEE Std 1003.2;
// 			if >=199309L, add IEEE Std 1003.1b-1993;
// 			if >=199506L, add IEEE Std 1003.1c-1995;
// 			if >=200112L, all of IEEE 1003.1-2004
// 			if >=200809L, all of IEEE 1003.1-2008
//    _XOPEN_SOURCE	Includes POSIX and XPG things.  Set to 500 if
// 			Single Unix conformance is wanted, to 600 for the
// 			sixth revision, to 700 for the seventh revision.
//    _XOPEN_SOURCE_EXTENDED XPG things and X/Open Unix extensions.
//    _LARGEFILE_SOURCE	Some more functions for correct standard I/O.
//    _LARGEFILE64_SOURCE	Additional functionality from LFS for large files.
//    _FILE_OFFSET_BITS=N	Select default filesystem interface.
//    _ATFILE_SOURCE	Additional *at interfaces.
//    _GNU_SOURCE		All of the above, plus GNU extensions.
//    _DEFAULT_SOURCE	The default set of features (taking precedence over
// 			__STRICT_ANSI__).
//
//    _FORTIFY_SOURCE	Add security hardening to many library functions.
// 			Set to 1 or 2; 2 performs stricter checks than 1.
//
//    _REENTRANT, _THREAD_SAFE
// 			Obsolete; equivalent to _POSIX_C_SOURCE=199506L.
//
//    The `-ansi' switch to the GNU C compiler, and standards conformance
//    options such as `-std=c99', define __STRICT_ANSI__.  If none of
//    these are defined, or if _DEFAULT_SOURCE is defined, the default is
//    to have _POSIX_SOURCE set to one and _POSIX_C_SOURCE set to
//    200809L, as well as enabling miscellaneous functions from BSD and
//    SVID.  If more than one of these are defined, they accumulate.  For
//    example __STRICT_ANSI__, _POSIX_SOURCE and _POSIX_C_SOURCE together
//    give you ISO C, 1003.1, and 1003.2, but nothing else.
//
//    These are defined by this file and are used by the
//    header files to decide what to declare or define:
//
//    __GLIBC_USE (F)	Define things from feature set F.  This is defined
// 			to 1 or 0; the subsequent macros are either defined
// 			or undefined, and those tests should be moved to
// 			__GLIBC_USE.
//    __USE_ISOC11		Define ISO C11 things.
//    __USE_ISOC99		Define ISO C99 things.
//    __USE_ISOC95		Define ISO C90 AMD1 (C95) things.
//    __USE_ISOCXX11	Define ISO C++11 things.
//    __USE_POSIX		Define IEEE Std 1003.1 things.
//    __USE_POSIX2		Define IEEE Std 1003.2 things.
//    __USE_POSIX199309	Define IEEE Std 1003.1, and .1b things.
//    __USE_POSIX199506	Define IEEE Std 1003.1, .1b, .1c and .1i things.
//    __USE_XOPEN		Define XPG things.
//    __USE_XOPEN_EXTENDED	Define X/Open Unix things.
//    __USE_UNIX98		Define Single Unix V2 things.
//    __USE_XOPEN2K        Define XPG6 things.
//    __USE_XOPEN2KXSI     Define XPG6 XSI things.
//    __USE_XOPEN2K8       Define XPG7 things.
//    __USE_XOPEN2K8XSI    Define XPG7 XSI things.
//    __USE_LARGEFILE	Define correct standard I/O things.
//    __USE_LARGEFILE64	Define LFS things with separate names.
//    __USE_FILE_OFFSET64	Define 64bit interface as default.
//    __USE_MISC		Define things from 4.3BSD or System V Unix.
//    __USE_ATFILE		Define *at interfaces and AT_* constants for them.
//    __USE_GNU		Define GNU extensions.
//    __USE_FORTIFY_LEVEL	Additional security measures used, according to level.
//
//    The macros `__GNU_LIBRARY__', `__GLIBC__', and `__GLIBC_MINOR__' are
//    defined by this file unconditionally.  `__GNU_LIBRARY__' is provided
//    only for compatibility.  All new code should use the other symbols
//    to test for features.
//
//    All macros listed above as possibly being defined by this file are
//    explicitly undefined if they are not explicitly defined.
//    Feature-test macros that are not defined by the user or compiler
//    but are implied by the other feature-test macros defined (or by the
//    lack of any definitions) are defined by the file.
//
//    ISO C feature test macros depend on the definition of the macro
//    when an affected header is included, not when the first system
//    header is included, and so they are handled in
//    <bits/libc-header-start.h>, which does not have a multiple include
//    guard.  Feature test macros that can be handled from the first
//    system header included are handled here.

// Undefine everything, so we get a clean slate.

// Suppress kernel-name space pollution unless user expressedly asks
//    for it.

// Convenience macro to test the version of gcc.
//    Use like this:
//    #if __GNUC_PREREQ (2,8)
//    ... code requiring gcc 2.8 or later ...
//    #endif
//    Note: only works for GCC 2.0 and later, because __GNUC_MINOR__ was
//    added in 2.0.

// Similarly for clang.  Features added to GCC after version 4.2 may
//    or may not also be available in clang, and clang's definitions of
//    __GNUC(_MINOR)__ are fixed at 4 and 2 respectively.  Not all such
//    features can be queried via __has_extension/__has_feature.

// Whether to use feature set F.

// _BSD_SOURCE and _SVID_SOURCE are deprecated aliases for
//    _DEFAULT_SOURCE.  If _DEFAULT_SOURCE is present we do not
//    issue a warning; the expectation is that the source is being
//    transitioned to use the new macro.

// If _GNU_SOURCE was defined by the user, turn on all the other features.

// If nothing (other than _GNU_SOURCE and _DEFAULT_SOURCE) is defined,
//    define _DEFAULT_SOURCE.

// This is to enable the ISO C11 extension.

// This is to enable the ISO C99 extension.

// This is to enable the ISO C90 Amendment 1:1995 extension.

// If none of the ANSI/POSIX macros are defined, or if _DEFAULT_SOURCE
//    is defined, use POSIX.1-2008 (or another version depending on
//    _XOPEN_SOURCE).

// Some C libraries once required _REENTRANT and/or _THREAD_SAFE to be
//    defined in all multithreaded code.  GNU libc has not required this
//    for many years.  We now treat them as compatibility synonyms for
//    _POSIX_C_SOURCE=199506L, which is the earliest level of POSIX with
//    comprehensive support for multithreaded code.  Using them never
//    lowers the selected level of POSIX conformance, only raises it.

// The function 'gets' existed in C89, but is impossible to use
//    safely.  It has been removed from ISO C11 and ISO C++14.  Note: for
//    compatibility with various implementations of <cstdio>, this test
//    must consider only the value of __cplusplus when compiling C++.

// Get definitions of __STDC_* predefined macros, if the compiler has
//    not preincluded this header automatically.
// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// This macro indicates that the installed library is the GNU C Library.
//    For historic reasons the value now is 6 and this will stay from now
//    on.  The use of this variable is deprecated.  Use __GLIBC__ and
//    __GLIBC_MINOR__ now (see below) when you want to test for a specific
//    GNU C library version and use the values in <gnu/lib-names.h> to get
//    the sonames of the shared libraries.

// Major and minor version number of the GNU C library package.  Use
//    these macros to test for features in specific releases.

// This is here only because every header file already includes this one.
// Copyright (C) 1992-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// We are almost always included from features.h.

// The GNU libc does not support any K&R compilers or the traditional mode
//    of ISO C compilers anymore.  Check for some of the combinations not
//    anymore supported.

// Some user header file might have defined this before.

// All functions, except those with callbacks or those that
//    synchronize memory, are leaf functions.

// GCC can always grok prototypes.  For C++ programs we add throw()
//    to help it optimize the function calls.  But this works only with
//    gcc 2.8.x and egcs.  For gcc 3.2 and up we even mark C functions
//    as non-throwing using a function attribute since programs can use
//    the -fexceptions options for C code as well.

// Compilers that are not clang may object to
//        #if defined __clang__ && __has_extension(...)
//    even though they do not need to evaluate the right-hand side of the &&.

// These two macros are not used in glibc anymore.  They are kept here
//    only because some other projects expect the macros to be defined.

// For these things, GCC behaves the ANSI way normally,
//    and the non-ANSI way under -traditional.

// This is not a typedef so `const __ptr_t' does the right thing.

// C++ needs to know that types and declarations are C, not C++.

// Fortify support.

// Support for flexible arrays.
//    Headers that should use flexible arrays only if they're "real"
//    (e.g. only if they won't affect sizeof()) should test
//    #if __glibc_c99_flexarr_available.

// __asm__ ("xyz") is used throughout the headers to rename functions
//    at the assembly language level.  This is wrapped by the __REDIRECT
//    macro, in order to support compilers that can do this some other
//    way.  When compilers don't support asm-names at all, we have to do
//    preprocessor tricks instead (which don't have exactly the right
//    semantics, but it's the best we can do).
//
//    Example:
//    int __REDIRECT(setpgrp, (__pid_t pid, __pid_t pgrp), setpgid);

//
// #elif __SOME_OTHER_COMPILER__
//
// # define __REDIRECT(name, proto, alias) name proto; 	_Pragma("let " #name " = " #alias)

// GCC has various useful declarations that can be made with the
//    `__attribute__' syntax.  All of the ways we use this do fine if
//    they are omitted for compilers that don't understand it.

// At some point during the gcc 2.96 development the `malloc' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// Tell the compiler which arguments to an allocation function
//    indicate the size of the allocation.

// At some point during the gcc 2.96 development the `pure' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// This declaration tells the compiler that the value is constant.

// At some point during the gcc 3.1 development the `used' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.

// Since version 3.2, gcc allows marking deprecated functions.

// Since version 4.5, gcc also allows one to specify the message printed
//    when a deprecated function is used.  clang claims to be gcc 4.2, but
//    may also support this feature.

// At some point during the gcc 2.8 development the `format_arg' attribute
//    for functions was introduced.  We don't want to use it unconditionally
//    (although this would be possible) since it generates warnings.
//    If several `format_arg' attributes are given for the same function, in
//    gcc-3.0 and older, all but the last one are ignored.  In newer gccs,
//    all designated arguments are considered.

// At some point during the gcc 2.97 development the `strfmon' format
//    attribute for functions was introduced.  We don't want to use it
//    unconditionally (although this would be possible) since it
//    generates warnings.

// The nonull function attribute allows to mark pointer parameters which
//    must not be NULL.

// If fortification mode, we warn about unused results of certain
//    function calls which can lead to problems.

// Forces a function to be always inlined.
// The Linux kernel defines __always_inline in stddef.h (283d7573), and
//    it conflicts with this definition.  Therefore undefine it first to
//    allow either header to be included first.

// Associate error messages with the source location of the call site rather
//    than with the source location inside the function.

// GCC 4.3 and above with -std=c99 or -std=gnu99 implements ISO C99
//    inline semantics, unless -fgnu89-inline is used.  Using __GNUC_STDC_INLINE__
//    or __GNUC_GNU_INLINE is not a good enough check for gcc because gcc versions
//    older than 4.3 may define these macros and still not guarantee GNU inlining
//    semantics.
//
//    clang++ identifies itself as gcc-4.2, but has support for GNU inlining
//    semantics, that can be checked fot by using the __GNUC_STDC_INLINE_ and
//    __GNUC_GNU_INLINE__ macro definitions.

// GCC 4.3 and above allow passing all anonymous arguments of an
//    __extern_always_inline function to some other vararg function.

// It is possible to compile containing GCC extensions even if GCC is
//    run in pedantic mode if the uses are carefully marked using the
//    `__extension__' keyword.  But this is not generally available before
//    version 2.8.

// __restrict is known in EGCS 1.2 and above.

// ISO C99 also allows to declare arrays as non-overlapping.  The syntax is
//      array_name[restrict]
//    GCC 3.1 supports this.

// Describes a char array whose address can safely be passed as the first
//    argument to strncpy and strncat, as the char array is not necessarily
//    a NUL-terminated string.

// Determine the wordsize from the preprocessor defines.

// Both x86-64 and x32 use the 64-bit system call interface.
// Properties of long double type.  ldbl-96 version.
//    Copyright (C) 2016-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License  published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// long double is distinct from double, so there is nothing to
//    define here.

// __glibc_macro_warning (MESSAGE) issues warning MESSAGE.  This is
//    intended for use in preprocessor macros.
//
//    Note: MESSAGE must be a _single_ string; concatenation of string
//    literals is not supported.

// Generic selection (ISO C11) is a C-only feature, available in GCC
//    since version 4.9.  Previous versions do not provide generic
//    selection, even though they might set __STDC_VERSION__ to 201112L,
//    when in -std=c11 mode.  Thus, we must check for !defined __GNUC__
//    when testing __STDC_VERSION__ for generic selection support.
//    On the other hand, Clang also defines __GNUC__, so a clang-specific
//    check is required to enable the use of generic selection.

// If we don't have __REDIRECT, prototypes will be missing if
//    __USE_FILE_OFFSET64 but not __USE_LARGEFILE[64].

// Decide whether we can define 'extern inline' functions in headers.

// This is here only because every header file already includes this one.
//    Get the definitions of all the appropriate `__stub_FUNCTION' symbols.
//    <gnu/stubs.h> contains `#define __stub_FUNCTION' when FUNCTION is a stub
//    that will always return failure (and set errno to ENOSYS).
// This file is automatically generated.
//    This file selects the right generated file of `__stub_FUNCTION' macros
//    based on the architecture being compiled for.

// This file is automatically generated.
//    It defines a symbol `__stub_FUNCTION' for each function
//    in the C library which is a stub, meaning it will fail
//    every time called, usually setting errno to ENOSYS.

// ISO/IEC TR 24731-2:2010 defines the __STDC_WANT_LIB_EXT2__
//    macro.

// ISO/IEC TS 18661-1:2014 defines the __STDC_WANT_IEC_60559_BFP_EXT__
//    macro.

// ISO/IEC TS 18661-4:2015 defines the
//    __STDC_WANT_IEC_60559_FUNCS_EXT__ macro.

// ISO/IEC TS 18661-3:2015 defines the
//    __STDC_WANT_IEC_60559_TYPES_EXT__ macro.

// Maximum length of any multibyte character in any locale.
//    We define this value here since the gcc header does not define
//    the correct value.

// If we are not using GNU CC we have to define all the symbols ourself.
//    Otherwise use gcc's definitions (see below).

// Get the compiler's limits.h, which defines almost all the ISO constants.
//
//     We put this #include_next outside the double inclusion check because
//     it should be possible to include this file more than once and still get
//     the definitions from gcc's header.

// The <limits.h> files in some gcc versions don't define LLONG_MIN,
//    LLONG_MAX, and ULLONG_MAX.  Instead only the values gcc defined for
//    ages are available.

// The integer width macros are not defined by GCC's <limits.h> before
//    GCC 7, or if _GNU_SOURCE rather than
//    __STDC_WANT_IEC_60559_BFP_EXT__ is used to enable this feature.

// POSIX adds things to <limits.h>.
// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

//	POSIX Standard: 2.9.2 Minimum Values	Added to <limits.h>
//
//	Never include this file directly; use <limits.h> instead.

// Determine the wordsize from the preprocessor defines.

// Both x86-64 and x32 use the 64-bit system call interface.

// These are the standard-mandated minimum values.

// Minimum number of operations in one list I/O call.

// Minimal number of outstanding asynchronous I/O operations.

// Maximum length of arguments to `execve', including environment.

// Maximum simultaneous processes per real user ID.

// Minimal number of timer expiration overruns.

// Maximum length of a host name (not including the terminating null)
//    as returned from the GETHOSTNAME function.

// Maximum link count of a file.

// Maximum length of login name.

// Number of bytes in a terminal canonical input queue.

// Number of bytes for which space will be
//    available in a terminal input queue.

// Maximum number of message queues open for a process.

// Maximum number of supported message priorities.

// Number of bytes in a filename.

// Number of simultaneous supplementary group IDs per process.

// Number of files one process can have open at once.

// Number of bytes in a pathname.

// Number of bytes than can be written atomically to a pipe.

// The number of repeated occurrences of a BRE permitted by the
//    REGEXEC and REGCOMP functions when using the interval notation.

// Minimal number of realtime signals reserved for the application.

// Number of semaphores a process can have.

// Maximal value of a semaphore.

// Number of pending realtime signals.

// Largest value of a `ssize_t'.

// Number of streams a process can have open at once.

// The number of bytes in a symbolic link.

// The number of symbolic links that can be traversed in the
//    resolution of a pathname in the absence of a loop.

// Number of timer for a process.

// Maximum number of characters in a tty name.

// Maximum length of a timezone name (element of `tzname').

// Maximum clock resolution in nanoseconds.

// Get the implementation-specific values for the above.
// Minimum guaranteed maximum values for system limits.  Linux version.
//    Copyright (C) 1993-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public License as
//    published by the Free Software Foundation; either version 2.1 of the
//    License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; see the file COPYING.LIB.  If
//    not, see <http://www.gnu.org/licenses/>.

// The kernel header pollutes the namespace with the NR_OPEN symbol
//    and defines LINK_MAX although filesystems have different maxima.  A
//    similar thing is true for OPEN_MAX: the limit can be changed at
//    runtime and therefore the macro must not be defined.  Remove this
//    after including the header if necessary.

// The kernel sources contain a file with all the needed information.
// SPDX-License-Identifier: GPL-2.0 WITH Linux-syscall-note

// Have to remove NR_OPEN?
// Have to remove LINK_MAX?
// Have to remove OPEN_MAX?
// Have to remove ARG_MAX?

// The number of data keys per process.
// This is the value this implementation supports.

// Controlling the iterations of destructors for thread-specific data.
// Number of iterations this implementation does.

// The number of threads per process.
// We have no predefined limit on the number of threads.

// Maximum amount by which a process can descrease its asynchronous I/O
//    priority level.

// Minimum size for a thread.  We are free to choose a reasonable value.

// Maximum number of timer expiration overruns.

// Maximum tty name length.

// Maximum login name length.  This is arbitrary.

// Maximum host name length.

// Maximum message queue priority level.

// Maximum value the semaphore can have.

// ssize_t is not formally required to be the signed type
//    corresponding to size_t, but it is for all configurations supported
//    by glibc.

// This value is a guaranteed minimum maximum.
//    The current maximum can be got from `sysconf'.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; include <limits.h> instead.

// The maximum `ibase' and `obase' values allowed by the `bc' utility.

// The maximum number of elements allowed in an array by the `bc' utility.

// The maximum `scale' value allowed by the `bc' utility.

// The maximum length of a string constant accepted by the `bc' utility.

// The maximum number of weights that can be assigned to an entry of
//    the LC_COLLATE `order' keyword in the locale definition file.

// The maximum number of expressions that can be nested
//    within parentheses by the `expr' utility.

// The maximum length, in bytes, of an input line.

// The maximum number of repeated occurrences of a regular expression
//    permitted when using the interval notation `\{M,N\}'.

// The maximum number of bytes in a character class name.  We have no
//    fixed limit, 2048 is a high number.

// These values are implementation-specific,
//    and may vary within the implementation.
//    Their precise values can be obtained from sysconf.

// This value is defined like this in regex.h.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free
// Software Foundation; either version 3, or (at your option) any later
// version.
//
// GCC is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License
// for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// Number of bits in a `char'.

// Maximum length of a multibyte character.

// Minimum and maximum values a `signed char' can hold.

// Maximum value an `unsigned char' can hold.  (Minimum is 0).

// Minimum and maximum values a `char' can hold.

// Minimum and maximum values a `signed short int' can hold.

// Maximum value an `unsigned short int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed int' can hold.

// Maximum value an `unsigned int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed long int' can hold.
//    (Same as `int').

// Maximum value an `unsigned long int' can hold.  (Minimum is 0).

// Minimum and maximum values a `signed long long int' can hold.

// Maximum value an `unsigned long long int' can hold.  (Minimum is 0).

// This administrivia gets added to the end of limits.h
//    if the system has its own version of limits.h.

var _ int8 /* gen.c:2:13: */
