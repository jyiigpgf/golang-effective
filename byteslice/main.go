package main

import "fmt"

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	*p = slice
	return len(data), nil
}

func main() {
	var b ByteSlice
	b = make([]byte, 500)
	i, err := fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(i, err)
	fmt.Println(b)
}
