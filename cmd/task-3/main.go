package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 2)

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
