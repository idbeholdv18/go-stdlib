package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	showDir := flag.Bool("dir", false, "set to show dirs")
	path := flag.String("path", ".", "dir path")

	flag.Parse()

	entries, err := os.ReadDir(*path)
	if err != nil {
		panic(err)
	}

	for _, e := range entries {
		if e.IsDir() && !*showDir {
			continue
		}
		fmt.Println(e)
	}
}
