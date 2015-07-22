package main

import "fmt"

func preIncr(x *int, y int) int {
	*x += y
	return *x
}

func postIncr(x *int, y int) int {
	xPre := *x
	*x += y
	return xPre
}

func main() {
	i := 0
	fmt.Printf("%v\n", postIncr(&i, 1))
	fmt.Printf("%v\n", preIncr(&i, 1))
}
