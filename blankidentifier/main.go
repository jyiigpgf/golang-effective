package main

import (
	"fmt"
	"io"
	"os"
)



var _ = fmt.Printf
//var myPrint = fmt.Println // For debugging; delete when done.
//var myPrint = func(...interface{}) {
//
//}
var _  io.Reader    // For debugging; delete when done.

type Puller interface {
	Pull()
	//Pull2()
}

type Pusher interface {
	Push()
}

type Github struct {

}

func (g *Github)Pull()  {
	fmt.Println("github pull")
}

var _ Puller = (*Github)(nil)
//var _ Pusher = (*Github)(nil)

func main() {
	//[1]
	fd, err := os.Open("test.go")
	if err != nil {
		//log.Fatal(err)
	}
	// TODO: use fd.
	_ = fd

	//myPrint("test")

	//[2]
	g := &Github{}
	g.Pull()
}