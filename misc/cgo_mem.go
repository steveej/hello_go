package main

// #include <stdlib.h>
// #include <string.h>
// #include <stdio.h>
//
// void
// new_string_by_ref(char **sref)
// {
//   *sref = strdup("some new bytes");
// }
// char*
// new_string_by_ret()
// {
//   return strdup("some new bytes");
// }
// char*
// new_string_by_malloc()
// {
//   char* new_string = "some new bytes";
//   int size = sizeof(char) * strlen(new_string) + 1;
//   char* dest_string = malloc(size);
//   memcpy(dest_string, new_string, size);
//   return dest_string;
// }
import "C"

import (
	"fmt"
	"github.com/cloudfoundry/gosigar"
	"os"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

func looseMem1(cleanup bool) (s string) {
	var pointer *C.char
	C.new_string_by_ref(&pointer)
	s = C.GoString(pointer)
	if cleanup {
		C.free(unsafe.Pointer(pointer))
	}
	return
}

func looseMem2(cleanup bool) (s string) {
	pointer := C.new_string_by_ret()
	s = C.GoString(pointer)
	if cleanup {
		C.free(unsafe.Pointer(pointer))
	}
	return
}

func looseMem3(cleanup bool) (s string) {
	pointer := C.new_string_by_malloc()
	s = C.GoString(pointer)
	if cleanup {
		C.free(unsafe.Pointer(pointer))
	}
	return
}

func getMem(pid int) (procMem sigar.ProcMem) {
	procMem = sigar.ProcMem{}
	procMem.Get(pid)
	return procMem
}

func getMemMiB(pid int) (procMem sigar.ProcMem) {
	procMem = getMem(pid)
	procMem.Size /= 1024 * 10e3
	procMem.Resident /= 1024 * 10e3
	return procMem
}

var funcs = []func(bool) string{
	looseMem1,
	looseMem2,
	looseMem3,
}

func funcName(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return nil
}

func main() {
	pid := os.Getpid()
	cleanup := true
	for _, f := range funcs {
		fmt.Printf("=== Testing function %v\n", funcName(f))
		t0 := time.Now()
		t := t0
		mem0 := getMem(pid)
		mem := mem0
		for i := 1; i < 10e6; i++ {
			f(cleanup)
			if t1 := time.Now(); t1.Sub(t).Seconds() > 1 {
				mem := getMem(pid)
				t = t1
				fmt.Printf("Size: %v Resident: %v\n", mem.Size, mem.Resident)
			}
		}
		mem = getMem(pid)
		fmt.Printf("Lost Size: %v Resident: %v\n", mem.Size-mem0.Size, mem.Resident-mem0.Resident)
	}
}
