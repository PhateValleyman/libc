// Code generated by 'ccgo time/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o time/time_linux_arm64.go -pkgname time', DO NOT EDIT.

package time

import (
	"math"
	"reflect"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ unsafe.Pointer

const (
	CLOCK_BOOTTIME            = 7
	CLOCK_BOOTTIME_ALARM      = 9
	CLOCK_MONOTONIC           = 1
	CLOCK_MONOTONIC_COARSE    = 6
	CLOCK_MONOTONIC_RAW       = 4
	CLOCK_PROCESS_CPUTIME_ID  = 2
	CLOCK_REALTIME            = 0
	CLOCK_REALTIME_ALARM      = 8
	CLOCK_REALTIME_COARSE     = 5
	CLOCK_TAI                 = 11
	CLOCK_THREAD_CPUTIME_ID   = 3
	TIMER_ABSTIME             = 1
	TIME_UTC                  = 1
	X_ATFILE_SOURCE           = 1
	X_BITS_TIME_H             = 1
	X_BITS_TYPESIZES_H        = 1
	X_BITS_TYPES_H            = 1
	X_BITS_TYPES_LOCALE_T_H   = 1
	X_BITS_TYPES___LOCALE_T_H = 1
	X_BSD_SIZE_T_             = 0
	X_BSD_SIZE_T_DEFINED_     = 0
	X_DEFAULT_SOURCE          = 1
	X_FEATURES_H              = 1
	X_FILE_OFFSET_BITS        = 64
	X_GCC_SIZE_T              = 0
	X_LP64                    = 1
	X_POSIX_C_SOURCE          = 200809
	X_POSIX_SOURCE            = 1
	X_SIZET_                  = 0
	X_SIZE_T                  = 0
	X_SIZE_T_                 = 0
	X_SIZE_T_DECLARED         = 0
	X_SIZE_T_DEFINED          = 0
	X_SIZE_T_DEFINED_         = 0
	X_STDC_PREDEF_H           = 1
	X_STRUCT_TIMESPEC         = 1
	X_SYS_CDEFS_H             = 1
	X_SYS_SIZE_T_H            = 0
	X_TIME_H                  = 1
	X_T_SIZE                  = 0
	X_T_SIZE_                 = 0
	Linux                     = 1
	Unix                      = 1
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = uint32 /* <builtin>:15:24 */

type X__int128_t = [2]int64   /* <builtin>:21:24 */ //TODO
type X__uint128_t = [2]uint64 /* <builtin>:22:25 */ //TODO

type X__builtin_va_list = uintptr /* <builtin>:44:14 */
type X__float128 = float64        /* <builtin>:45:21 */

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

//  In 4.3bsd-net2, leave these undefined to indicate that size_t, etc.
//     are already defined.
//  BSD/OS 3.1 and FreeBSD [23].x require the MACHINE_ANSI_H check here.
//  NetBSD 5 requires the I386_ANSI_H and X86_64_ANSI_H checks here.

// A null pointer constant.

// This defines CLOCKS_PER_SEC, which is the number of processor clock
//    ticks per second, and possibly a number of other constants.
// System-dependent timing definitions.  Linux version.
//    Copyright (C) 1996-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <time.h> instead.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

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

// Determine the wordsize from the preprocessor defines.
//
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

// Convenience types.
type X__u_char = uint8   /* types.h:30:23 */
type X__u_short = uint16 /* types.h:31:28 */
type X__u_int = uint32   /* types.h:32:22 */
type X__u_long = uint64  /* types.h:33:27 */

// Fixed-size types, underlying types depend on word size and compiler.
type X__int8_t = int8     /* types.h:36:21 */
type X__uint8_t = uint8   /* types.h:37:23 */
type X__int16_t = int16   /* types.h:38:26 */
type X__uint16_t = uint16 /* types.h:39:28 */
type X__int32_t = int32   /* types.h:40:20 */
type X__uint32_t = uint32 /* types.h:41:22 */
type X__int64_t = int64   /* types.h:43:25 */
type X__uint64_t = uint64 /* types.h:44:27 */

// Smallest types with at least a given width.
type X__int_least8_t = X__int8_t     /* types.h:51:18 */
type X__uint_least8_t = X__uint8_t   /* types.h:52:19 */
type X__int_least16_t = X__int16_t   /* types.h:53:19 */
type X__uint_least16_t = X__uint16_t /* types.h:54:20 */
type X__int_least32_t = X__int32_t   /* types.h:55:19 */
type X__uint_least32_t = X__uint32_t /* types.h:56:20 */
type X__int_least64_t = X__int64_t   /* types.h:57:19 */
type X__uint_least64_t = X__uint64_t /* types.h:58:20 */

// quad_t is also 64 bits.
type X__quad_t = int64    /* types.h:62:18 */
type X__u_quad_t = uint64 /* types.h:63:27 */

// Largest integral types.
type X__intmax_t = int64   /* types.h:71:18 */
type X__uintmax_t = uint64 /* types.h:72:27 */

// The machine-dependent file <bits/typesizes.h> defines __*_T_TYPE
//    macros for each of the OS types we define below.  The definitions
//    of those macros must use the following macros for underlying types.
//    We define __S<SIZE>_TYPE and __U<SIZE>_TYPE for the signed and unsigned
//    variants of each of the following integer types on this machine.
//
// 	16		-- "natural" 16-bit type (always short)
// 	32		-- "natural" 32-bit type (always int)
// 	64		-- "natural" 64-bit type (long or long long)
// 	LONG32		-- 32-bit type, traditionally long
// 	QUAD		-- 64-bit type, always long long
// 	WORD		-- natural type of __WORDSIZE bits (int or long)
// 	LONGWORD	-- type of __WORDSIZE bits, traditionally long
//
//    We distinguish WORD/LONGWORD, 32/LONG32, and 64/QUAD so that the
//    conventional uses of `long' or `long long' type modifiers match the
//    types we define, even when a less-adorned type would be the same size.
//    This matters for (somewhat) portably writing printf/scanf formats for
//    these types, where using the appropriate l or ll format modifiers can
//    make the typedefs and the formats match up across all GNU platforms.  If
//    we used `long' when it's 64 bits where `long long' is expected, then the
//    compiler would warn about the formats not matching the argument types,
//    and the programmer changing them to shut up the compiler would break the
//    program's portability.
//
//    Here we assume what is presently the case in all the GCC configurations
//    we support: long long is always 64 bits, long is always word/address size,
//    and int is always 32 bits.

// No need to mark the typedef with __extension__.
// bits/typesizes.h -- underlying types for *_t.  For the generic Linux ABI.
//    Copyright (C) 2011-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//    Contributed by Chris Metcalf <cmetcalf@tilera.com>, 2011.
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
//    License along with the GNU C Library.  If not, see
//    <http://www.gnu.org/licenses/>.

// See <bits/types.h> for the meaning of these macros.  This file exists so
//    that <bits/types.h> need not vary across different GNU platforms.

// Tell the libc code that off_t and off64_t are actually the same type
//    for all ABI purposes, even if possibly expressed as different base types
//    for C type-checking purposes.

// Same for ino_t and ino64_t.

// And for __rlim_t and __rlim64_t.

// Number of descriptors that can fit in an `fd_set'.

type X__dev_t = uint64                     /* types.h:143:25 */ // Type of device numbers.
type X__uid_t = uint32                     /* types.h:144:25 */ // Type of user identifications.
type X__gid_t = uint32                     /* types.h:145:25 */ // Type of group identifications.
type X__ino_t = uint64                     /* types.h:146:25 */ // Type of file serial numbers.
type X__ino64_t = uint64                   /* types.h:147:27 */ // Type of file serial numbers (LFS).
type X__mode_t = uint32                    /* types.h:148:26 */ // Type of file attribute bitmasks.
type X__nlink_t = uint32                   /* types.h:149:27 */ // Type of file link counts.
type X__off_t = int64                      /* types.h:150:25 */ // Type of file sizes and offsets.
type X__off64_t = int64                    /* types.h:151:27 */ // Type of file sizes and offsets (LFS).
type X__pid_t = int32                      /* types.h:152:25 */ // Type of process identifications.
type X__fsid_t = struct{ F__val [2]int32 } /* types.h:153:26 */ // Type of file system IDs.
type X__clock_t = int64                    /* types.h:154:27 */ // Type of CPU usage counts.
type X__rlim_t = uint64                    /* types.h:155:26 */ // Type for resource measurement.
type X__rlim64_t = uint64                  /* types.h:156:28 */ // Type for resource measurement (LFS).
type X__id_t = uint32                      /* types.h:157:24 */ // General type for IDs.
type X__time_t = int64                     /* types.h:158:26 */ // Seconds since the Epoch.
type X__useconds_t = uint32                /* types.h:159:30 */ // Count of microseconds.
type X__suseconds_t = int64                /* types.h:160:31 */ // Signed count of microseconds.

type X__daddr_t = int32 /* types.h:162:27 */ // The type of a disk address.
type X__key_t = int32   /* types.h:163:25 */ // Type of an IPC key.

// Clock ID used in clock and timer functions.
type X__clockid_t = int32 /* types.h:166:29 */

// Timer ID returned by `timer_create'.
type X__timer_t = uintptr /* types.h:169:12 */

// Type to represent block size.
type X__blksize_t = int32 /* types.h:172:29 */

// Types from the Large File Support interface.

// Type to count number of disk blocks.
type X__blkcnt_t = int64   /* types.h:177:28 */
type X__blkcnt64_t = int64 /* types.h:178:30 */

// Type to count file system blocks.
type X__fsblkcnt_t = uint64   /* types.h:181:30 */
type X__fsblkcnt64_t = uint64 /* types.h:182:32 */

// Type to count file system nodes.
type X__fsfilcnt_t = uint64   /* types.h:185:30 */
type X__fsfilcnt64_t = uint64 /* types.h:186:32 */

// Type of miscellaneous file system fields.
type X__fsword_t = int64 /* types.h:189:28 */

type X__ssize_t = int64 /* types.h:191:27 */ // Type of a byte count, or error.

// Signed long type used in system calls.
type X__syscall_slong_t = int64 /* types.h:194:33 */
// Unsigned long type used in system calls.
type X__syscall_ulong_t = uint64 /* types.h:196:33 */

// These few don't really vary by system, they always correspond
//    to one of the other defined types.
type X__loff_t = X__off64_t /* types.h:200:19 */ // Type of file sizes and offsets (LFS).
type X__caddr_t = uintptr   /* types.h:201:14 */

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int64 /* types.h:204:25 */

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* types.h:207:23 */

// C99: An integer type that can be accessed as an atomic entity,
//    even in the presence of asynchronous interrupts.
//    It is not currently necessary for this to be machine-specific.
type X__sig_atomic_t = int32 /* types.h:212:13 */

// ISO/IEC 9899:1999 7.23.1: Components of time
//    The macro `CLOCKS_PER_SEC' is an expression with type `clock_t' that is
//    the number per second of the value returned by the `clock' function.
// CAE XSH, Issue 4, Version 2: <time.h>
//    The value of CLOCKS_PER_SEC is required to be 1 million on all
//    XSI-conformant systems.

// Identifier for system-wide realtime clock.
// Monotonic system-wide clock.
// High-resolution timer from the CPU.
// Thread-specific CPU-time clock.
// Monotonic system-wide clock, not adjusted for frequency scaling.
// Identifier for system-wide realtime clock, updated only on ticks.
// Monotonic system-wide clock, updated only on ticks.
// Monotonic system-wide clock that includes time spent in suspension.
// Like CLOCK_REALTIME but also wakes suspended system.
// Like CLOCK_BOOTTIME but also wakes suspended system.
// Like CLOCK_REALTIME but in International Atomic Time.

// Flag to indicate time is absolute.

// Many of the typedefs and structs whose official home is this header
//    may also need to be defined by other headers.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// Returned by `clock'.
type Clock_t = X__clock_t /* clock_t.h:7:19 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// Returned by `time'.
type Time_t = X__time_t /* time_t.h:7:18 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// ISO C `broken-down time' structure.
type Tm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	_          [4]byte
	Ftm_gmtoff int64
	Ftm_zone   uintptr
} /* struct_tm.h:7:1 */

// NB: Include guard matches what <linux/time.h> uses.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// POSIX.1b structure for a time value.  This is like a `struct timeval' but
//    has nanoseconds instead of microseconds.
type Timespec = struct {
	Ftv_sec  X__time_t
	Ftv_nsec X__syscall_slong_t
} /* struct_timespec.h:9:1 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// Clock ID used in clock and timer functions.
type Clockid_t = X__clockid_t /* clockid_t.h:7:21 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// Timer ID returned by `timer_create'.
type Timer_t = X__timer_t /* timer_t.h:7:19 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
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

// Never include this file directly; use <sys/types.h> instead.

// NB: Include guard matches what <linux/time.h> uses.

// POSIX.1b structure for timer start values and intervals.
type Itimerspec = struct {
	Fit_interval struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
	Fit_value struct {
		Ftv_sec  X__time_t
		Ftv_nsec X__syscall_slong_t
	}
} /* struct_itimerspec.h:8:1 */

type Pid_t = X__pid_t /* time.h:54:17 */

// Definition of locale_t.
//    Copyright (C) 2017-2018 Free Software Foundation, Inc.
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

// Definition of struct __locale_struct and __locale_t.
//    Copyright (C) 1997-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//    Contributed by Ulrich Drepper <drepper@cygnus.com>, 1997.
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

// POSIX.1-2008: the locale_t type, representing a locale context
//    (implementation-namespace version).  This type should be treated
//    as opaque by applications; some details are exposed for the sake of
//    efficiency in e.g. ctype functions.

type X__locale_struct = struct {
	F__locales       [13]uintptr
	F__ctype_b       uintptr
	F__ctype_tolower uintptr
	F__ctype_toupper uintptr
	F__names         [13]uintptr
} /* __locale_t.h:28:1 */

type X__locale_t = uintptr /* __locale_t.h:42:32 */

type Locale_t = X__locale_t /* locale_t.h:24:20 */

var _ int8 /* gen.c:2:13: */
