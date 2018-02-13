package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

func main() {
	done := make(chan struct{})
	runtime.LockOSThread()

	go func() {

		defer func() {
			close(done)
			syscall.Syscall(syscall.SYS_EXIT, uintptr(0), 0, 0)
		}()

		_, _, e := syscall.Syscall(syscall.SYS_SETGID, uintptr(1000), 0, 0)
		if e != 0 {
			fmt.Fprintf(os.Stderr, "setgid() error: %v\n", e.Error())
			return
		}

		_, _, e = syscall.Syscall(syscall.SYS_SETUID, uintptr(65534), 0, 0)
		if e != 0 {
			fmt.Fprintf(os.Stderr, "setuid() error: %v\n", e.Error())
			return
		}
	}()

	<-done

	tcpFile, err := os.Open("/proc/net/tcp")
	if err != nil {
		fmt.Fprintf(os.Stderr, "open() error: %v\n", err)
		return
	}

	fmt.Printf("Alles jut: %v\n", tcpFile != nil)
}
