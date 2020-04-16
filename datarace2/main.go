package main

import (
	"time"
)

func main() {
	 t := 65535

	 go func() {
	 	for {
			//time.Sleep(time.Second)
	 		//fmt.Println(t)
			t = t + 1
		}
	 }()

	go func() {
		for {
			//time.Sleep(time.Second)
			//fmt.Println(t)
			t = t + 1
		}
	}()

	 go func() {
	 	for {
	 		//time.Sleep(time.Second)
			t = t + 1
		}
	 }()

	 time.Sleep(time.Hour)
}
