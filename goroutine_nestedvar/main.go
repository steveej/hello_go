package main

import (
	"fmt"
	"sync"
	"time"
)

type Request struct {
	number int
	sleep  int
	n      *int
}

func serve(queue chan Request) {
	var (
		wgInner sync.WaitGroup
	)
	for req := range queue {
		wgInner.Add(1)
		realreq := req
		f := func() {
			myn := &req.n
			realn := **myn

			wgInner.Done()
			time.Sleep(time.Duration(req.sleep) * time.Second)
			fmt.Println(req.number, realreq.number, myn, *myn, **myn, &realn, realn)
		}
		go f()
		wgInner.Wait()
	}
}

func main() {
	queue := make(chan Request)
	go serve(queue)

	max := 10
	for n := 0; n < max; n++ {
		queue <- Request{n, n, &n}
	}
	time.Sleep(time.Duration(max) * time.Second)
}
