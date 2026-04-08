package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll("test_dir")
	if err != nil {
		log.Fatal(err)
	}
}
