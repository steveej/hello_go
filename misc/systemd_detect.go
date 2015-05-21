package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <sys/types.h>
// #include <stdlib.h>
// #include <errno.h>
//
// int
// my_sd_pid_get_owner_uid(void *f, pid_t pid, uid_t *uid)
// {
//   int (*sd_pid_get_owner_uid)(pid_t, uid_t *);
//
//   sd_pid_get_owner_uid = (int (*)(pid_t, uid_t *))f;
//   return sd_pid_get_owner_uid(pid, uid);
// }
//
// int
// my_sd_pid_get_unit(void *f, pid_t pid, char **unit)
// {
//   int (*sd_pid_get_unit)(pid_t, char **);
//
//   sd_pid_get_unit = (int (*)(pid_t, char **))f;
//   return sd_pid_get_unit(pid, unit);
// }
//
// int
// my_sd_pid_get_user_unit(void *f, pid_t pid, char **user_unit)
// {
//   int (*sd_pid_get_user_unit)(pid_t, char **);
//
//   sd_pid_get_user_unit = (int (*)(pid_t, char **))f;
//   return sd_pid_get_user_unit(pid, user_unit);
// }
//
// int
// my_sd_pid_get_session(void *f, pid_t pid, char **session)
// {
//   int (*sd_pid_get_session)(pid_t, char **);
//
//   sd_pid_get_session = (int (*)(pid_t, char **))f;
//   return sd_pid_get_session(pid, session);
// }
//
// int
// my_sd_pid_get_slice(void *f, pid_t pid, char **slice)
// {
//   int (*sd_pid_get_slice)(pid_t, char **);
//
//   sd_pid_get_slice = (int (*)(pid_t, char **))f;
//   return sd_pid_get_slice(pid, slice);
// }
//
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

func runningFromUnitFile() (ret bool, err error) {
	handle := C.dlopen(C.CString("libsystemd.so"), C.RTLD_LAZY)
	if handle == nil {
		// we can't open libsystemd-login.so so we assume systemd is not
		// installed and we're not running from a unit file
		fmt.Printf("libsystemd.so not found")
		ret = false
		return
	}
	defer func() {
		if r := C.dlclose(handle); r != 0 {
			err = fmt.Errorf("error closing libsystemd-login.so")
		}
	}()

	sd_pid_get_owner_uid := C.dlsym(handle, C.CString("sd_pid_get_owner_uid"))
	if sd_pid_get_owner_uid == nil {
		err = fmt.Errorf("error resolving sd_pid_get_owner_uid function")
		return
	}
	var uid C.uid_t
	uid_errno := C.my_sd_pid_get_owner_uid(sd_pid_get_owner_uid, 0, &uid)
	fmt.Printf("sd_pid_get_owner_uid: %v, %v\n", uid_errno, uid)

	sd_pid_get_unit := C.dlsym(handle, C.CString("sd_pid_get_unit"))
	if sd_pid_get_unit == nil {
		err = fmt.Errorf("error resolving sd_pid_get_unit function")
		return
	}
	var unit *C.char
	defer C.free(unsafe.Pointer(unit))
	unit_errno := C.my_sd_pid_get_unit(sd_pid_get_unit, 0, &unit)
	fmt.Printf("sd_pid_get_unit: %v, %v\n", unit_errno, C.GoString(unit))

	sd_pid_get_user_unit := C.dlsym(handle, C.CString("sd_pid_get_user_unit"))
	if sd_pid_get_user_unit == nil {
		err = fmt.Errorf("error resolving sd_pid_get_user_unit function")
		return
	}
	var user_unit *C.char
	defer C.free(unsafe.Pointer(user_unit))
	user_unit_errno := C.my_sd_pid_get_user_unit(sd_pid_get_user_unit, 0, &user_unit)
	fmt.Printf("sd_pid_get_user_unit: %v, %v\n", user_unit_errno, C.GoString(user_unit))

	sd_pid_get_session := C.dlsym(handle, C.CString("sd_pid_get_session"))
	if sd_pid_get_session == nil {
		err = fmt.Errorf("error resolving sd_pid_get_session function")
		return
	}
	var session *C.char
	defer C.free(unsafe.Pointer(session))
	session_errno := C.my_sd_pid_get_session(sd_pid_get_session, 0, &session)
	fmt.Printf("sd_pid_get_session: %v, %v\n", session_errno, C.GoString(session))

	sd_pid_get_slice := C.dlsym(handle, C.CString("sd_pid_get_slice"))
	if sd_pid_get_slice == nil {
		err = fmt.Errorf("error resolving sd_pid_get_slice function")
		return
	}
	var slice *C.char
	defer C.free(unsafe.Pointer(slice))
	slice_errno := C.my_sd_pid_get_slice(sd_pid_get_slice, 0, &slice)
	fmt.Printf("sd_pid_get_slice: %v, %v\n", slice_errno, C.GoString(slice))

	ret = C.GoString(slice) == "system.service"
	return
}

func main() {
	insideUnit, err := runningFromUnitFile()

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("inside unit? %v\n", insideUnit)
		os.Exit(0)
	}
}
