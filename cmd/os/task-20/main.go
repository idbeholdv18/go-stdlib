package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("you must be specify src and dst")
	}
	srcFilepath := os.Args[1]
	dstFilepath := os.Args[2]

	srcFile, err := os.Open(srcFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	dstFile, err := os.Create(dstFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dstFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	
	if err := dstFile.Chmod(srcInfo.Mode()); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		log.Fatal(err)
	}
}
