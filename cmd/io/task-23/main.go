package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	a := flag.Bool("a", false, "find all files by filename")
	flag.Parse()

	filename := flag.Arg(0)
	if filename == "" {
		log.Fatal("filename must be specified")
	}

	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filename == d.Name() {
			fmt.Println(path)
			if !*a {
				return filepath.SkipAll
			}
		}
		return nil
	})
}
