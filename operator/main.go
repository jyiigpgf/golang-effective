package main

import "fmt"

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}

func main() {
	s := Sequence{}
	fmt.Println(len(s))
}