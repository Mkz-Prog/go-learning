package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		if 'a' <= b[i] && b[i] <= 'z' {
			if b[i] <= 'm' {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		} else if 'A' <= b[i] && b[i] <= 'Z' {
			if b[i] <= 'M' {
				b[i] += 13
			} else {
				b[i] -= 13
			}
		}
	}
	return n, err
}

// Exercise: rot13Reader from a Tour of Go
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
