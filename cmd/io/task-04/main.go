package main

import (
	"io"
	"log"
	"os"
)

type teeReader struct {
	r io.Reader
	w io.Writer
}

func (t *teeReader) Read(buf []byte) (n int, err error) {
	n, err = t.r.Read(buf)
	if n > 0 {
		if n2, err2 := t.w.Write(buf[:n]); err2 != nil {
			return n2, err2
		}
	}
	return
}

func TeeReader(r io.Reader, w io.Writer) io.Reader {
	return &teeReader{
		r: r,
		w: w,
	}
}

func main() {
	file, err := os.Open("test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dst, err := os.Create("test2.log")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	r := TeeReader(file, dst)

	if _, err := io.Copy(io.Discard, r); err != nil {
		log.Fatal(err)
	}
}
