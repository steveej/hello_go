package main

import (
	"flag"
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

var pattern, dir string

func init() {
	const (
		dirDefault     = "."
		dirHelp        = "Directory to list"
		patternDefault = "^.*\\.(\\?\\!swp).*$"
		patternHelp    = "Regular expression pattern that will be used to filter the listed files."
	)
	flag.StringVar(&pattern, "pattern", patternDefault, patternHelp)
	flag.StringVar(&dir, "dir", dirDefault, dirHelp)
}

func main() {
	flag.Parse()

	files, err := listFiles(dir)
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
