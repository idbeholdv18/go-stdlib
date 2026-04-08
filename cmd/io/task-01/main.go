package main

import (
	"fmt"
)

type MyReader struct{}

func (r *MyReader) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	r := &MyReader{}

	buf := make([]byte, 4)

	for i := 0; i < 5; i++ {
		n, _ := r.Read(buf)
		fmt.Printf("%s", buf[:n])
	}
	fmt.Println()
}
