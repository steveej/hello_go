package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	stat := syscall.Statfs_t{}
	err := syscall.Statfs(os.Args[1], &stat)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//bsize := stat.Bsize
	//fmt.Println(stat)
	s := fmt.Sprintf(`
    Statfs_t {
        Type    %d
        Bsize   %d
        Blocks  %d
        Bfree   %d
        Bavail  %d
        Files   %d
        Ffree   %d
        Frsize  %d
        Flags   %d
    }
	`, stat.Type,
		stat.Bsize,
		stat.Blocks,
		stat.Bfree,
		stat.Bavail,
		stat.Files,
		stat.Ffree,
		stat.Frsize,
		stat.Flags)

	fmt.Println(s)
}
