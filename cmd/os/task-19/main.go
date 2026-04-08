package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename must be specified")
	}
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 4096)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Stdout.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
}
