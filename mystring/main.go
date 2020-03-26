package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type MyString string

func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", string(m))
}

type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := time.Time(t).Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func main() {
	s := MyString("ttttt")
	fmt.Println(s)


	t := MyTime(time.Now())
	b, _ :=json.Marshal(t)
	fmt.Println(string(b))

	tInterface := interface{}(t)
	switch tInterface.(type) {
	case time.Time:
		fmt.Println("time.Time")
	case MyTime:
		fmt.Println("MyTime")

	}
}