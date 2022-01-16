//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"modernc.org/cc/v3"
	ccgo "modernc.org/ccgo/v3/lib"
)

var (
	goos   = runtime.GOOS
	goarch = runtime.GOARCH
)

func origin(skip int) string {
	pc, fn, fl, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	var fns string
	if f != nil {
		fns = f.Name()
		if x := strings.LastIndex(fns, "."); x > 0 {
			fns = fns[x+1:]
		}
	}
	return fmt.Sprintf("%s:%d:%s", filepath.Base(fn), fl, fns)
}

func trc(s string, args ...interface{}) string { //TODO-
	switch {
	case s == "":
		s = fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)
	default:
		s = fmt.Sprintf(s, args...)
	}
	_, fn, fl, _ := runtime.Caller(1)
	r := fmt.Sprintf("\n%s:%d: TRC %s", fn, fl, s)
	fmt.Fprintf(os.Stdout, "%s\n", r)
	os.Stdout.Sync()
	return r
}

func main() {
	fmt.Printf("Running on %s/%s.\n", runtime.GOOS, runtime.GOARCH)
	if s := os.Getenv("TARGET_GOOS"); s != "" {
		goos = s
	}
	if s := os.Getenv("TARGET_GOARCH"); s != "" {
		goarch = s
	}
	g := []string{"libc.go", "mem.go"}
	switch goos {
	case "linux":
		g = append(g, "libc_unix.go", "pthread.go", "pthread_all.go")
		makeMuslLinux(goos, goarch)
	case "freebsd":
		g = append(g, "libc_unix.go", "pthread.go", "pthread_all.go")
		makeMuslFreeBSD(goos, goarch)
	case "netbsd":
		g = append(g, "libc_unix.go", "pthread.go", "pthread_all.go")
		makeMuslNetBSD(goos, goarch)
	case "openbsd":
		g = append(g, "libc_unix.go", "pthread.go", "pthread_all.go")
		makeMuslFreeBSD(goos, goarch)
	case "darwin":
		g = append(g, "libc_unix.go", "pthread.go", "pthread_all.go")
		makeMuslDarwin(goos, goarch)
	case "windows":
		g = append(g, "libc_windows.go")
		makeMuslWin(goos, goarch)
	}
	_, _, hostSysIncludes, err := cc.HostConfig(os.Getenv("CCGO_CPP"))
	if err != nil {
		fail(err)
	}

	x, err := filepath.Glob(fmt.Sprintf("*_%s.go", goos))
	if err != nil {
		fail(err)
	}

	g = append(g, x...)
	if x, err = filepath.Glob(fmt.Sprintf("*_%s_%s.go", goos, goarch)); err != nil {
		fail(err)
	}

	g = append(g, x...)
	m := map[string]struct{}{}
	for _, v := range g {
		f, err := os.Open(v)
		if err != nil {
			fail(err)
		}

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			s := sc.Text()
			s0 := s
			switch {
			case strings.HasPrefix(s, "func X"):
				s = s[len("func X"):]
				x := strings.IndexByte(s, '(')
				s = s[:x]
			case strings.HasPrefix(s, "var X"):
				s = s[len("var X"):]
				x := strings.IndexByte(s, ' ')
				s = s[:x]
			default:
				continue
			}

			if s == "" {
				panic(fmt.Sprintf("internal error %q", s0))
			}
			m[s] = struct{}{}
		}
		if err := sc.Err(); err != nil {
			fail(err)
		}

		f.Close()
	}
	var a []string
	for k := range m {
		a = append(a, k)
	}
	sort.Strings(a)
	b := bytes.NewBuffer(nil)
	b.WriteString(`// Code generated by 'go generate' - DO NOT EDIT.

package libc // import "modernc.org/libc"

var CAPI = map[string]struct{}{`)

	for _, v := range a {
		fmt.Fprintf(b, "\n\t%q: {},", v)
	}
	b.WriteString("\n}")
	if err := ioutil.WriteFile(fmt.Sprintf("capi_%s_%s.go", goos, goarch), b.Bytes(), 0660); err != nil {
		fail(err)
	}

	ccgoHelpers()

	if err := libcHeaders(hostSysIncludes); err != nil {
		fail(err)
	}
}

func makeMuslWin(goos, goarch string) {
	if runtime.GOOS != "windows" {
		switch goarch {
		case "386":
			makeMuslWin386(goos, goarch)
			return
		default:
			fail(fmt.Errorf("must be runned on Windows"))
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	case "386":
		arch = "i386"
	case "arm":
		arch = "arm"
	case "arm64":
		arch = "aarch64"
	default:
		fail(fmt.Errorf("unknown/unsupported GOARCH: %q", goarch))
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	//TODO- run("mkdir", "-p", "obj/include/bits")
	//TODO- run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	//TODO- run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	//TODO- run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))
	if err := os.MkdirAll("obj\\include\\bits", 0770); err != nil {
		fail(err)
	}

	run("cmd", "/c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("cmd", "/c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("cmd", "/c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))

	if _, err := runcc(
		"-D__environ=environ",
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		// Keep the below lines sorted.
		"src/ctype/isalnum.c",
		"src/ctype/isalpha.c",
		"src/ctype/isdigit.c",
		"src/ctype/islower.c",
		"src/ctype/isprint.c",
		"src/ctype/isspace.c",
		"src/ctype/isxdigit.c",
		"src/env/putenv.c",
		"src/env/setenv.c",
		"src/env/unsetenv.c",
		"src/multibyte/wcrtomb.c",
		"src/multibyte/wcsrtombs.c",
		"src/multibyte/wcstombs.c",
		"src/string/strchrnul.c",
		"src/string/strdup.c",
	); err != nil {
		fail(err)
	}
}

func makeMuslWin386(goos, goarch string) {
	if runtime.GOOS != "linux" {
		fail(fmt.Errorf("must be runned on Linux"))
	}

	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	arch := "i386"
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	run("mkdir", "-p", "obj/include/bits")
	run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))

	if _, err := runcc(
		"-D__environ=environ",
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		// Keep the below lines sorted.
		"src/ctype/isalnum.c",
		"src/ctype/isalpha.c",
		"src/ctype/isdigit.c",
		"src/ctype/islower.c",
		"src/ctype/isprint.c",
		"src/ctype/isspace.c",
		"src/ctype/isxdigit.c",
		"src/env/putenv.c",
		"src/env/setenv.c",
		"src/env/unsetenv.c",
		"src/multibyte/wcrtomb.c",
		"src/multibyte/wcsrtombs.c",
		"src/multibyte/wcstombs.c",
		"src/string/strchrnul.c",
		"src/string/strdup.c",
	); err != nil {
		fail(err)
	}
}

func makeMuslDarwin(goos, goarch string) {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	case "arm64":
		arch = "aarch64"
	default:
		fail(fmt.Errorf("unknown/unsupported GOARCH: %q", goarch))
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	run("mkdir", "-p", "obj/include/bits")
	run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))
	if _, err := runcc(
		"-D__environ=environ",
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6",
		"-hide", "isascii,isspace,tolower,toupper",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		"../darwin/table.c",

		// Keep the below lines sorted.
		"src/env/putenv.c",
		"src/env/setenv.c",
		"src/env/unsetenv.c",
		"src/internal/floatscan.c",
		"src/internal/intscan.c",
		"src/internal/shgetc.c",
		"src/locale/localeconv.c",
		"src/math/__fpclassify.c",
		"src/math/__fpclassifyf.c",
		"src/math/__fpclassifyl.c",
		"src/math/copysignl.c",
		"src/math/fabsl.c",
		"src/math/fmodl.c",
		"src/math/nanf.c",
		"src/math/rint.c",
		"src/math/scalbn.c",
		"src/math/scalbnl.c",
		"src/network/freeaddrinfo.c",
		"src/network/getaddrinfo.c",
		"src/network/gethostbyaddr.c",
		"src/network/gethostbyaddr_r.c",
		"src/network/gethostbyname.c",
		"src/network/gethostbyname2.c",
		"src/network/gethostbyname2_r.c",
		"src/network/getnameinfo.c",
		"src/network/h_errno.c",
		"src/network/inet_aton.c",
		"src/network/inet_ntop.c",
		"src/network/inet_pton.c",
		"src/network/lookup_ipliteral.c",
		"src/network/lookup_name.c",
		"src/network/lookup_serv.c",
		"src/prng/rand_r.c",
		"src/stdio/__toread.c",
		"src/stdio/__uflow.c",
		"src/stdlib/strtod.c",
		"src/stdlib/strtol.c",
		"src/string/strchrnul.c",
		"src/string/strdup.c",
		"src/string/strlcat.c",
		"src/string/strlcpy.c",
		"src/string/strncasecmp.c",
		"src/string/strncat.c",
		"src/string/strnlen.c",
		"src/string/strspn.c",
		"src/string/strtok.c",
	); err != nil {
		fail(err)
	}
}

func makeMuslLinux(goos, goarch string) {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	case "386":
		arch = "i386"
	case "arm":
		arch = "arm"
	case "arm64":
		arch = "aarch64"
	case "s390x":
		arch = "s390x"
	default:
		fail(fmt.Errorf("unknown/unsupported GOARCH: %q", goarch))
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	run("mkdir", "-p", "obj/include/bits")
	run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))
	if _, err := runcc(
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		// Keep the below lines sorted.
		"src/ctype/__ctype_b_loc.c",
		"src/ctype/isalnum.c",
		"src/ctype/isalpha.c",
		"src/ctype/isdigit.c",
		"src/ctype/islower.c",
		"src/ctype/isprint.c",
		"src/ctype/isupper.c",
		"src/ctype/isxdigit.c",
		"src/dirent/closedir.c",
		"src/dirent/opendir.c",
		"src/dirent/readdir.c",
		"src/internal/floatscan.c",
		"src/internal/intscan.c",
		"src/internal/shgetc.c",
		"src/locale/localeconv.c",
		"src/math/__fpclassify.c",
		"src/math/__fpclassifyf.c",
		"src/math/__fpclassifyl.c",
		"src/math/copysignl.c",
		"src/math/fabsl.c",
		"src/math/fmodl.c",
		"src/math/nanf.c",
		"src/math/rint.c",
		"src/math/scalbn.c",
		"src/math/scalbnl.c",
		"src/multibyte/internal.c",
		"src/multibyte/mbrtowc.c",
		"src/multibyte/mbsinit.c",
		"src/network/freeaddrinfo.c",
		"src/network/getaddrinfo.c",
		"src/network/gethostbyaddr.c",
		"src/network/gethostbyaddr_r.c",
		"src/network/gethostbyname.c",
		"src/network/gethostbyname2.c",
		"src/network/gethostbyname2_r.c",
		"src/network/gethostbyname_r.c",
		"src/network/getnameinfo.c",
		"src/network/h_errno.c",
		"src/network/inet_aton.c",
		"src/network/inet_ntop.c",
		"src/network/inet_pton.c",
		"src/network/lookup_ipliteral.c",
		"src/network/lookup_name.c",
		"src/network/lookup_serv.c",
		"src/prng/rand_r.c",
		"src/stdio/__lockfile.c",
		"src/stdio/__toread.c",
		"src/stdio/__uflow.c",
		"src/stdio/sscanf.c",
		"src/stdio/vfscanf.c",
		"src/stdio/vsscanf.c",
		"src/stdlib/strtod.c",
		"src/stdlib/strtol.c",
		"src/string/strdup.c",
		"src/string/strlcat.c",
		"src/string/strlcpy.c",
		"src/string/strncasecmp.c",
		"src/string/strncat.c",
		"src/string/strnlen.c",
		"src/string/strspn.c",
		"src/string/strtok.c",
		"src/thread/pthread_attr_get.c",
		"src/thread/pthread_attr_setdetachstate.c",
		"src/thread/pthread_mutex_lock.c",
		"src/thread/pthread_mutexattr_destroy.c",
		"src/thread/pthread_mutexattr_init.c",
		"src/thread/pthread_mutexattr_settype.c",
	); err != nil {
		fail(err)
	}
}

func makeMuslNetBSD(goos, goarch string) {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	default:
		fail(fmt.Errorf("unknown/unsupported GOARCH: %q", goarch))
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	run("mkdir", "-p", "obj/include/bits")
	run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))
	if _, err := runcc(
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6,getnameinfo,gethostbyaddr_r,",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		"../netbsd/ctype_.cpp.c",

		// Keep the below lines sorted.
		"src/ctype/isalnum.c",
		"src/ctype/isalpha.c",
		"src/ctype/isdigit.c",
		"src/internal/floatscan.c",
		"src/internal/intscan.c",
		"src/internal/shgetc.c",
		"src/math/copysignl.c",
		"src/math/fabsl.c",
		"src/math/fmodl.c",
		"src/math/rint.c",
		"src/math/scalbn.c",
		"src/math/scalbnl.c",
		"src/network/freeaddrinfo.c",
		"src/network/getaddrinfo.c",
		"src/network/gethostbyaddr.c",
		"src/network/gethostbyaddr_r.c",
		"src/network/gethostbyname.c",
		"src/network/gethostbyname2.c",
		"src/network/gethostbyname2_r.c",
		"src/network/getnameinfo.c",
		"src/network/h_errno.c",
		"src/network/inet_aton.c",
		"src/network/inet_ntop.c",
		"src/network/inet_pton.c",
		"src/network/lookup_ipliteral.c",
		"src/network/lookup_name.c",
		"src/network/lookup_serv.c",
		"src/stdio/__toread.c",
		"src/stdio/__uflow.c",
		"src/stdlib/strtod.c",
		"src/stdlib/strtol.c",
		"src/string/strdup.c",
		"src/string/strnlen.c",
		"src/string/strspn.c",
	); err != nil {
		fail(err)
	}
}

func makeMuslFreeBSD(goos, goarch string) {
	wd, err := os.Getwd()
	if err != nil {
		fail(err)
	}

	if err := os.Chdir("musl"); err != nil {
		fail(err)
	}

	var arch string
	switch goarch {
	case "amd64":
		arch = "x86_64"
	case "386":
		arch = "i386"
	default:
		fail(fmt.Errorf("unknown/unsupported GOARCH: %q", goarch))
	}
	defer func() {
		if err := os.Chdir(wd); err != nil {
			fail(err)
		}
	}()

	run("mkdir", "-p", "obj/include/bits")
	run("sh", "-c", fmt.Sprintf("sed -f ./tools/mkalltypes.sed ./arch/%s/bits/alltypes.h.in ./include/alltypes.h.in > obj/include/bits/alltypes.h", arch))
	run("sh", "-c", fmt.Sprintf("cp arch/%s/bits/syscall.h.in obj/include/bits/syscall.h", arch))
	run("sh", "-c", fmt.Sprintf("sed -n -e s/__NR_/SYS_/p < arch/%s/bits/syscall.h.in >> obj/include/bits/syscall.h", arch))
	if _, err := runcc(
		"-export-externs", "X",
		"-hide", "__syscall0,__syscall1,__syscall2,__syscall3,__syscall4,__syscall5,__syscall6,getnameinfo,gethostbyaddr_r,",
		"-nostdinc",
		"-nostdlib",
		"-o", fmt.Sprintf("../musl_%s_%s.go", goos, goarch),
		"-pkgname", "libc",
		"-static-locals-prefix", "_s",

		// Keep the order below, don't sort!
		fmt.Sprintf("-I%s", filepath.Join("arch", arch)),
		fmt.Sprintf("-I%s", "arch/generic"),
		fmt.Sprintf("-I%s", "obj/src/internal"),
		fmt.Sprintf("-I%s", "src/include"),
		fmt.Sprintf("-I%s", "src/internal"),
		fmt.Sprintf("-I%s", "obj/include"),
		fmt.Sprintf("-I%s", "include"),
		// Keep the order above, don't sort!

		"copyright.c", // Inject legalese first

		"../freebsd/table.cpp.c",

		// Keep the below lines sorted.
		"src/ctype/isalnum.c",
		"src/ctype/isalpha.c",
		"src/ctype/isdigit.c",
		"src/ctype/isprint.c",
		"src/ctype/isspace.c",
		"src/internal/floatscan.c",
		"src/internal/intscan.c",
		"src/internal/shgetc.c",
		"src/math/copysignl.c",
		"src/math/fabsl.c",
		"src/math/fmodl.c",
		"src/math/rint.c",
		"src/math/scalbn.c",
		"src/math/scalbnl.c",
		"src/network/freeaddrinfo.c",
		"src/network/getaddrinfo.c",
		"src/network/gethostbyaddr.c",
		"src/network/gethostbyaddr_r.c",
		"src/network/gethostbyname.c",
		"src/network/gethostbyname2.c",
		"src/network/gethostbyname2_r.c",
		"src/network/getnameinfo.c",
		"src/network/h_errno.c",
		"src/network/inet_aton.c",
		"src/network/inet_ntop.c",
		"src/network/inet_pton.c",
		"src/network/lookup_ipliteral.c",
		"src/network/lookup_name.c",
		"src/network/lookup_serv.c",
		"src/stdio/__toread.c",
		"src/stdio/__uflow.c",
		"src/stdlib/strtod.c",
		"src/stdlib/strtol.c",
		"src/string/strdup.c",
		"src/string/strnlen.c",
		"src/string/strspn.c",
	); err != nil {
		fail(err)
	}
}

func run(arg0 string, args ...string) []byte {
	fmt.Printf("%s %q\n", arg0, args)
	cmd := exec.Command(arg0, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		sout := strings.TrimSpace(string(out))
		fmt.Fprintf(os.Stderr, "==== FAIL\n%s\n%s\n", sout, err)
		fail(err)
	}
	return out
}

type echoWriter struct {
	w bytes.Buffer
}

func (w *echoWriter) Write(b []byte) (int, error) {
	os.Stdout.Write(b)
	return w.w.Write(b)
}

func runcc(args ...string) ([]byte, error) {
	args = append([]string{"ccgo"}, args...)
	// fmt.Printf("%q\n", args)
	var out echoWriter
	err := ccgo.NewTask(args, &out, &out).Main()
	return out.w.Bytes(), err
}

func libcHeaders(paths []string) error {
	const cfile = "gen.c"
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		path = filepath.Clean(path)
		if strings.HasPrefix(path, ".") {
			return nil
		}

		dir := path
		ok := false
		for _, v := range paths {
			full := filepath.Join(v, dir+".h")
			if fi, err := os.Stat(full); err == nil && !fi.IsDir() {
				ok = true
				break
			}
		}
		if !ok {
			return nil
		}

		var src string
		switch filepath.ToSlash(path) {
		case "fts":
			src = `
#include <sys/types.h>
#include <sys/stat.h>
#include <fts.h>
`
		default:
			src = fmt.Sprintf("#include <%s.h>\n", dir)
		}
		src += "static char _;\n"
		fn := filepath.Join(dir, cfile)
		if err := ioutil.WriteFile(fn, []byte(src), 0660); err != nil {
			return err
		}

		defer os.Remove(fn)

		dest := filepath.Join(path, fmt.Sprintf("%s_%s_%s.go", filepath.Base(path), goos, goarch))
		base := filepath.Base(dir)
		argv := []string{
			fn,

			"-crt-import-path", "",
			"-export-defines", "",
			"-export-enums", "",
			"-export-externs", "X",
			"-export-fields", "F",
			"-export-structs", "",
			"-export-typedefs", "",
			"-header",
			"-hide", "_OSSwapInt16,_OSSwapInt32,_OSSwapInt64",
			"-o", dest,
			"-pkgname", base,
		}
		out, err := runcc(argv...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s%s\n", path, out, err)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n%s", path, out)
		}
		return nil
	})
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func ccgoHelpers() {
	var (
		signed = []string{
			"int8",
			"int16",
			"int32",
			"int64",
		}
		unsigned = []string{
			"uint8",
			"uint16",
			"uint32",
			"uint64",
		}
		ints   = append(signed[:len(signed):len(signed)], unsigned...)
		intptr = append(ints[:len(ints):len(ints)], "uintptr")
		arith  = append(ints[:len(ints):len(ints)], "float32", "float64")
		scalar = append(arith[:len(arith):len(arith)], []string{"uintptr"}...)
		sizes  = []string{"8", "16", "32", "64"}
		atomic = []string{
			"int32",
			"int64",
			"uint32",
			"uint64",
			"uintptr",
		}
	)

	b := bytes.NewBuffer(nil)
	b.WriteString(`// Code generated by 'go generate' - DO NOT EDIT.

package libc // import "modernc.org/libc"

import (
	"sync/atomic"
	"unsafe"
)

`)
	for _, v := range atomic {
		fmt.Fprintln(b)
		fmt.Fprintf(b, "func AtomicStoreN%s(ptr uintptr, val %s, memorder int32) { atomic.Store%[1]s((*%[2]s)(unsafe.Pointer(ptr)), val) }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range atomic {
		fmt.Fprintln(b)
		fmt.Fprintf(b, "func AtomicLoadN%s(ptr uintptr, memorder int32) %s { return atomic.Load%[1]s((*%[2]s)(unsafe.Pointer(ptr))) }\n", capitalize(v), v)
	}

	for _, v := range scalar {
		fmt.Fprintf(b, "func Assign%s(p *%s, v %[2]s) %[2]s { *p = v; return v }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) = v; return v }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignMul%s(p *%s, v %[2]s) %[2]s { *p *= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignDiv%s(p *%s, v %[2]s) %[2]s { *p /= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignRem%s(p *%s, v %[2]s) %[2]s { *p %%= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignAdd%s(p *%s, v %[2]s) %[2]s { *p += v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignSub%s(p *%s, v %[2]s) %[2]s { *p -= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignAnd%s(p *%s, v %[2]s) %[2]s { *p &= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignXor%s(p *%s, v %[2]s) %[2]s { *p ^= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignOr%s(p *%s, v %[2]s) %[2]s { *p |= v; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignMulPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) *= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignDivPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) /= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignRemPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) %%= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignAddPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) += v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func AssignSubPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) -= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignAndPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) &= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignXorPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) ^= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignOrPtr%s(p uintptr, v %s) %[2]s { *(*%[2]s)(unsafe.Pointer(p)) |= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShlPtr%s(p uintptr, v int) %s { *(*%[2]s)(unsafe.Pointer(p)) <<= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShrPtr%s(p uintptr, v int) %s { *(*%[2]s)(unsafe.Pointer(p)) >>= v; return *(*%[2]s)(unsafe.Pointer(p)) }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShl%s(p *%s, v int) %[2]s { *p <<= v; return *p }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func AssignShr%s(p *%s, v int) %[2]s { *p >>= v; return *p }\n\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PreInc%s(p *%s, d %[2]s) %[2]s { *p += d; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range atomic {
		fmt.Fprintf(b, "func PreIncAtomic%s(p *%s, d %[2]s) %[2]s { return atomic.Add%[1]s(p, d) }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PreDec%s(p *%s, d %[2]s) %[2]s { *p -= d; return *p }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range atomic {
		fmt.Fprintf(b, "func PreDecAtomic%s(p *%s, d %[2]s) %[2]s { return atomic.Add%[1]s(p, -d) }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PostInc%s(p *%s, d %[2]s) %[2]s { r := *p; *p += d; return r }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range atomic {
		fmt.Fprintf(b, "func PostIncAtomic%s(p *%s, d %[2]s) %[2]s { return atomic.Add%[1]s(p, d) - d }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func PostDec%s(p *%s, d %[2]s) %[2]s { r := *p; *p -= d; return r }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range atomic {
		fmt.Fprintf(b, "func PostDecAtomic%s(p *%s, d %[2]s) %[2]s { return atomic.Add%[1]s(p, -d) + d }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		for _, w := range scalar {
			fmt.Fprintf(b, "func %sFrom%s(n %s) %s { return %[4]s(n) }\n", capitalize(v), capitalize(w), w, v)
		}
	}

	fmt.Fprintln(b)
	for _, v := range scalar {
		fmt.Fprintf(b, "func %s(n %s) %[2]s { return n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func Neg%s(n %s) %[2]s { return -n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range intptr {
		fmt.Fprintf(b, "func Cpl%s(n %s) %[2]s { return ^n }\n", capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, v := range ints {
		fmt.Fprintf(b, `
func Bool%s(b bool) %s {
	if b {
		return 1
	}
	return 0
}
`, capitalize(v), v)
	}

	fmt.Fprintln(b)
	for _, sz := range sizes {
		for _, v := range ints {
			fmt.Fprintf(b, `
func SetBitFieldPtr%s%s(p uintptr, v %s, off int, mask uint%[1]s) {
	*(*uint%[1]s)(unsafe.Pointer(p)) = *(*uint%[1]s)(unsafe.Pointer(p))&^uint%[1]s(mask) | uint%[1]s(v)<<off&mask
}

`, sz, capitalize(v), v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func AssignBitFieldPtr%dInt%d(p uintptr, v int%[2]d, w, off int, mask uint%[1]d) int%[2]d {
	*(*uint%[1]d)(unsafe.Pointer(p)) = *(*uint%[1]d)(unsafe.Pointer(p))&^uint%[1]d(mask) | uint%[1]d(v)<<off&mask
	s := %[2]d - w
	return v << s >> s
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func AssignBitFieldPtr%dUint%d(p uintptr, v uint%[2]d, w, off int, mask uint%[1]d) uint%[2]d {
	*(*uint%[1]d)(unsafe.Pointer(p)) = *(*uint%[1]d)(unsafe.Pointer(p))&^uint%[1]d(mask) | uint%[1]d(v)<<off&mask
	return v & uint%[2]d(mask >> off)
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostDecBitFieldPtr%dInt%d(p uintptr, d int%[2]d, w, off int, mask uint%[1]d) (r int%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	s := %[2]d - w
	r = int%[2]d(x0) & int%[2]d(mask) << s >> (s+off)
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r-d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostDecBitFieldPtr%dUint%d(p uintptr, d uint%[2]d, w, off int, mask uint%[1]d) (r uint%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	r = uint%[2]d(x0) & uint%[2]d(mask) >> off
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r-d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostIncBitFieldPtr%dInt%d(p uintptr, d int%[2]d, w, off int, mask uint%[1]d) (r int%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	s := %[2]d - w
	r = int%[2]d(x0) & int%[2]d(mask) << s >> (s+off)
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r+d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	fmt.Fprintln(b)
	for _, sz := range []int{8, 16, 32, 64} {
		for _, v := range []int{8, 16, 32, 64} {
			fmt.Fprintf(b, `
func PostIncBitFieldPtr%dUint%d(p uintptr, d uint%[2]d, w, off int, mask uint%[1]d) (r uint%[2]d) {
	x0 := *(*uint%[1]d)(unsafe.Pointer(p))
	r = uint%[2]d(x0) & uint%[2]d(mask) >> off
	*(*uint%[1]d)(unsafe.Pointer(p)) = x0&^uint%[1]d(mask) | uint%[1]d(r+d)<<off&mask
	return r
}

`, sz, v)
		}
	}

	b.WriteString("\n")
	if err := ioutil.WriteFile(fmt.Sprintf("ccgo.go"), b.Bytes(), 0660); err != nil {
		fail(err)
	}
}

func capitalize(s string) string { return strings.ToUpper(s[:1]) + s[1:] }
