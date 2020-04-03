package main

import (
	"encoding/json"
	"fmt"
)

type Reader interface {
	Read() string
}

type WReader interface {
	Read() string
	Write(text string)
}

type Author struct {
	Text string
}

func (a Author)Read() string {
	return a.Text
}

func (a *Author)Write(text string) {
	a.Text = text
}

func (a Author) MarshalJSON() ([]byte, error) {
	return []byte("\"author MarshalJSON\""), nil
}

func (a *Author) UnmarshalJSON(b []byte) error {
	a.Text = "UnmarshalJSON"
	return nil
}


func main() {
	var a interface{}
	a = &Author{
		Text: "hello world",
	}
	switch a.(type) {
	case WReader:
	fmt.Println(a.(WReader).Read())
	a.(WReader).Write("my world")
	fmt.Println(a.(WReader).Read())
	case Reader:
		fmt.Println("reader")
	}

	data, err := json.Marshal(a)
	fmt.Println(err)
	fmt.Println(string(data))

	err = json.Unmarshal([]byte("\"\""), a)
	fmt.Println(err)
	fmt.Println(a.(WReader).Read())
}
