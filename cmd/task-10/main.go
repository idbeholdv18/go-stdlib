package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("filename must be specified")
	}

	root, err := os.OpenRoot(".")
	if err != nil {
		panic(err)
	}
	file, err := root.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Printf("%s", buf[:n])
	}
	fmt.Println()
}
