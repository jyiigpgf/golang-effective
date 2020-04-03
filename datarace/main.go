package main

import (
	"fmt"
	"time"
)

var MaxOutstanding uint = 1


var quit = make(chan bool, 0)
var sem = make(chan int, MaxOutstanding)

type Request struct {
	Text string
}

func process(req *Request) {
	fmt.Println(req.Text)
}

func Serve(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}

func main() {
	queue := make(chan *Request, 2)
	go Serve(queue)


	go func() {
		for {
			queue <- &Request{Text:"a"}
		}
	}()


	go func() {
		for {
			queue <- &Request{Text: "b"}
		}
	}()

	for {
		time.Sleep(5 * time.Second)
	}
}
