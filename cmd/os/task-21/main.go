package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	p := flag.Bool("p", false, "Create intermediate directories as required")
	flag.Parse()
	dirPath := flag.Arg(0)
	if dirPath == "" {
		log.Fatal("dir path must be specified")
	}

	if *p {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
