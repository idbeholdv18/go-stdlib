package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	a := flag.Bool("a", false, "list all files")
	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") && !*a {
			continue
		}
		if file.IsDir() {
			fmt.Printf("d\t%6s\t%s\n", "", file.Name())
		} else {
			info, err := file.Info()
			if err != nil {
				fmt.Printf("error reading file info for %s: %v\n", file.Name(), err)
				continue
			}
			fmt.Printf("-\t%6d\t%s\n", info.Size(), file.Name())
		}
	}
}
