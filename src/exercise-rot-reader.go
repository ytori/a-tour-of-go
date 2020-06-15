package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return (c-'A'+13)%26 + 'A'
	case 'a' <= c && c <= 'z':
		return (c-'a'+13)%26 + 'a'
	default:
		return c
	}
}

func (reader rot13Reader) Read(p []byte) (int, error) {

	n, err := reader.r.Read(p)
	if err != nil {
		return n, err
	}

	for i, v := range p {
		p[i] = rot13(v)
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
