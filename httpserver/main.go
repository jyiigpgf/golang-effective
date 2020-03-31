package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func MyHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v %v\n", w, req)
}



func main() {


	s := &http.Server{
		Addr:           ":8080",
		//Handler:        ctr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//ctr := new(Counter)
	//http.Handle("/", ctr)
	http.Handle("/", http.HandlerFunc(MyHandler))
	log.Fatal(s.ListenAndServe())

}