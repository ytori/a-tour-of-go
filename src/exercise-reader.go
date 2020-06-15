package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {

	for i, _ := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
