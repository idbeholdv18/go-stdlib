package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("specify filename")
	}
	path := os.Args[1]
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("failed to remove file %s: %v", path, err)
	}

	fmt.Printf("file %s has been removed\n", path)
}
