package main

import (
	"fmt"
	"github.com/hydrogen18/stoppableListener"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func f() {
	defer wg.Done()
	fmt.Println("f()")

	originalListener, err := net.Listen("tcp", ":12345")
	if err != nil {
		panic(err)
	}
	sl, err := stoppableListener.New(originalListener)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to %v", r.URL.Path)
		sl.Stop()
	})

	server := http.Server{}
	server.Serve(sl)
}

func g() {
	defer wg.Done()
	fmt.Println("g()")
	res, err := http.Get("http://localhost:12345/info")
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got: %s\n", text)
}

func main() {
	wg.Add(1)
	wg.Add(1)

	go f()
	go g()

	wg.Wait()
	fmt.Println("Main exiting.")
	os.Exit(0)
}
