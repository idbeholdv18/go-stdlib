package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Mkdir("test_dir", 0755); err != nil {
		log.Fatal(err)
	}
}
