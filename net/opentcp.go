package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	tcpFile, err := os.Open("/proc/net/tcp")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer tcpFile.Close()

	fmt.Println("Open TCP ports")
	scanner := bufio.NewScanner(tcpFile)
	for scanner.Scan() {
		re := regexp.MustCompile(`:([A-Z0-9]+) `)
		line := scanner.Text()
		result := re.FindAllStringSubmatch(line, -1)
		if result != nil {
			i, err := strconv.ParseInt(result[0][1], 16, 32)
			if err != nil {
				fmt.Printf("%v", err)
				return
			}
			fmt.Printf("- %d\n", i)
		}
	}
}
