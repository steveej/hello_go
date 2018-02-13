package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/opencontainers/runc/libcontainer/system"
)

func setUidGid(uid, gid int) error {
	fmt.Printf("Setting uid %d and gid %d\n", uid, gid)
	err := system.Setgid(gid)
	if err != nil {
		return fmt.Errorf("Setgid error: %v", err)
	}

	err = system.Setuid(uid)
	if err != nil {
		return fmt.Errorf("Setuid error: %v", err)
	}
	return nil
}

func printIds(i int) {
	fmt.Printf(
		"gorutine %d: uid=%d euid=%d gid=%d egid=%d\n", i,
		syscall.Getuid(), syscall.Geteuid(),
		syscall.Getgid(), syscall.Getegid(),
	)
}

func init() {
	runtime.LockOSThread()
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s UID GID\n", os.Args[0])
		os.Exit(1)
	}

	uid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Unknown UID:", err)
		os.Exit(1)
	}
	gid, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Unknown GID:", err)
		os.Exit(1)
	}

	for i := 1; i < 10; i++ {
		go func(i int) {
			err = setUidGid(uid, gid)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for {
				printIds(i)
				time.Sleep(5e8)
			}
		}(i)
	}

	for {
		runtime.Gosched()
	}
}
