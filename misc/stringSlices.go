package main

import (
	"fmt"
	"strings"
)

type StringSet struct {
	m map[string]struct{}
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]struct{})}
}

func (set *StringSet) Exists(s string) (existed bool) {
	_, existed = set.m[s]
	return
}

func (set *StringSet) Insert(s string) (existed bool) {
	existed = set.Exists(s)
	if !existed {
		set.m[s] = struct{}{}
	} else {
		existed = !existed
	}
	return
}

func main() {
	argstring := "s1,s2,s3,s4"

	splitStrings := strings.Split(argstring, ",")
	fmt.Println(splitStrings)

	stringSet := NewStringSet()

	for _, s := range splitStrings {
		fmt.Println(s)
		stringSet.Insert(s)
	}
	fmt.Println(stringSet)
	for k, v := range stringSet.m {
		fmt.Println(k, v)
	}
}
