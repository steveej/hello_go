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

func serve(queue chan *Request) {
	var (
		wgInner sync.WaitGroup
	)
	for req := range queue {
		wgInner.Add(1)
		realreq := req
		go func() {
			myn := &req.n
			realn := **myn

			wgInner.Done()
			time.Sleep(time.Duration(req.sleep) * time.Second)
			fmt.Println(req.number, realreq.number, myn, *myn, **myn, &realn, realn)
		}()
		wgInner.Wait()
	}
}

func main() {
	queue := make(chan *Request)
	go serve(queue)

	n := 1
	queue <- &Request{n, n, &n}

	n = 0
	queue <- &Request{n, n, &n}

	time.Sleep(2 * time.Second)
}
