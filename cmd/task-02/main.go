package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello, world!"))
	if err != nil {
		panic(err)
	}
}
