package main

import (
	"io"
	"os"
	"strings"
)

//var rot int = 'N' - 'A'

type Rot13Reader struct {
	r io.Reader
}

func (reader *Rot13Reader) Read(b []byte) (int, error) {
	var a, z, ua, uz byte = 'a', 'z', 'A', 'Z'
	var diff byte = a - ua

	n, e := reader.r.Read(b)
	for i, v := range b {
		var offset byte
		switch {
		case v <= uz && v >= ua: //if
			offset = 0
		case v <= z && v >= a: //else if
			offset, v = diff, v-diff
		default: //else
			break
		}

		if t := v + 13; t > uz {
			b[i] = t%uz + ua - 1
		} else {
			b[i] = t
		}
		b[i] += offset
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	//io.Copy  Writer , Reader
	io.Copy(os.Stdout, &r)
}
