package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func main() {
	var handler Hello
	err := http.ListenAndServe("localhost:4000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
