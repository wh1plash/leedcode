package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
	n, err := rot.r.Read(p)
	for i := 0; i < n; i++ {
		c := p[i]
		switch {
		case 'A' <= c && c <= 'Z':
			p[i] = 'A' + (c-'A'+13)%26
		case 'a' <= c && c <= 'z':
			p[i] = 'a' + (c-'a'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
