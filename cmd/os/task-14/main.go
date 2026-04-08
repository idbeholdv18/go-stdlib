package main

import (
	"fmt"
	"os"
)

func main() {
	filename := []byte("test.txt")
	file, err := os.OpenFile(string(filename), os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("file %s created\n", filename)
}
