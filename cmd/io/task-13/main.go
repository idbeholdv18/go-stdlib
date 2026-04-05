package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("some additional text"))
	if err != nil {
		panic(err)
	}
}
