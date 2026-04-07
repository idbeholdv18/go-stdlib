package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString("some text")
	if err != nil {
		panic(err)
	}
	file.Sync()
	file.Seek(0, 0)

	text, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(text))
}
