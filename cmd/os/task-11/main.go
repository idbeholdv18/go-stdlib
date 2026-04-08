package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("HOME env isn't set")
	}

	fmt.Println(home)
}
