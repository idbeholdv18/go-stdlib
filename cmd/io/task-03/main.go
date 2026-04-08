package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CountBytes(r io.Reader) (int, error) {
	n, err := io.Copy(io.Discard, r)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func main() {
	file, err := os.Open("test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	n, err := CountBytes(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}
