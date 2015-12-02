package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i, v := range b {
		b[i] = rot13(v)
	}
	return
}

func rot13(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case b >= 'a' && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	}
	return b
}
