package main

import (
	"fmt"
	"io"
	"strings"
)

//!!Read!!
func main() {
	//just liek StringReader
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) // []byte slice len(8) cap(8)

	for { // = for true (while)
		//func (T) Read(b []byte) (n int, err error)
		n, err := r.Read(b) // n how many byte read
		fmt.Printf("n = %v , b = %v (%p) , err =%v \n", n, b, b, err)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { // in java we use n < 0 to judge if we are in EOF now
			break
		}
	}
}
