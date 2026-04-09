package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type multiReader struct {
	readers []io.Reader
}

func (r *multiReader) Read(buf []byte) (int, error) {
	for len(r.readers) > 0 {
		n, err := r.readers[0].Read(buf)
		if err == io.EOF {
			r.readers = r.readers[1:]
			if n > 0 {
				return n, nil
			}
			continue
		}
		return n, err
	}
	return 0, io.EOF
}

func MultiReader(readers ...io.Reader) io.Reader {
	return &multiReader{
		readers: readers,
	}
}

func main() {
	r1 := strings.NewReader("hello")
	r2 := strings.NewReader(" world")

	r3 := MultiReader(r1, r2)

	res, err := io.ReadAll(r3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(res))
}
