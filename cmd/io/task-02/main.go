package main

import (
	"fmt"
	"io"
	"log"
)

type MyReader struct {
	buf    []byte
	offset int
}

func (r *MyReader) Read(buf []byte) (n int, err error) {
	if r.offset >= len(r.buf) {
		return 0, io.EOF
	}
	for i := range buf {
		if r.offset >= len(r.buf) {
			return i, nil
		}
		buf[i] = r.buf[r.offset]
		r.offset++
	}

	return len(buf), nil
}

func main() {
	r := &MyReader{
		buf:    []byte("hello world"),
		offset: 0,
	}

	buf := make([]byte, 2)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", string(buf[:n]))
	}
	fmt.Println()
}
