package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(str))
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (str *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s", str.Greeting, str.Punct, str.Who)
}

func main() {
	// your http.Handle calls here
	//在 web 服务器中注册它们来处理指定的路径。
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
