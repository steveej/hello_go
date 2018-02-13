package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	fmt.Printf("Error: %v\n", err)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
}
