package main

import "fmt"

func main() {
	c := 'b'
	switch c {
	default:
		fmt.Println("default")
	case 'a', 'b':
	case 'c':
		fmt.Println("b")

	}
}
