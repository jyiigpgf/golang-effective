package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 1)
	//ch2 := make(chan int, 2)
	ch <- 1
	go func() {
		t, ok := <- ch
		fmt.Println(t, ok)
	}()

	go func() {
		t, ok := <- ch
		fmt.Println(t, ok)
	}()

	go func() {
		t, ok := <- ch
		fmt.Println(t, ok)
	}()

	//a, b := <- ch , <- ch2
	//fmt.Println(a, b)
	//
	//switch t {
	//case 1, 2, 3:
	//	fmt.Println(t)
	//}
	//
	//close(ch)
	//t, ok = <- ch
	//fmt.Println(t, ok)


	time.Sleep(3 * time.Second)

	//contex, can := context.WithCancel(context.Background())
	//
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		t := <-contex.Done()
	//		fmt.Printf("done %p\n", &t)
	//	}
	//}()
	//time.Sleep(time.Second)
	//can()
	//time.Sleep(time.Second*100)

}
