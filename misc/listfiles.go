package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func listFiles(dir string) ([]string, error) {
	dirents, err := ioutil.ReadDir(dir)
	switch {
	case err == nil:
	case os.IsNotExist(err):
		return nil, nil
	default:
		return nil, err
	}

	files := []string{}
	for _, dent := range dirents {
		if dent.IsDir() {
			continue
		}

		files = append(files, dent.Name())
	}

	return files, nil
}

func filterFiles(strings []string, keepPattern string) ([]string, error) {
	fmt.Printf("Keeping files that match %v\n", keepPattern)
	new_strings := []string{}
	for _, s := range strings {
		match, err := regexp.MatchString(keepPattern, s)
		if err != nil {
			return nil, err
		} else if match == false {
			fmt.Printf("Filtering %v\n", s)
		} else {
			new_strings = append(new_strings, s)
		}
	}
	return new_strings, nil
}

func main() {
	pattern := "^.*\\.(?!swp).**$"
	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments.")
		os.Exit(1)
	} else if len(os.Args) == 3 {
		pattern = os.Args[2]
	}

	files, err := listFiles(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	files, err = filterFiles(files, pattern)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%v\n", files)
		os.Exit(0)
	}
}
