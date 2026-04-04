package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile, err := os.Open("./test.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("./test.copy.txt")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()

	// сделал намеренно для эксперимента. Лучше в данном случае использовать io.Copy
	buf := make([]byte, 2)
	_, err = io.CopyBuffer(dstFile, srcFile, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("file %s has been copied to %s\n", dstFile.Name(), srcFile.Name())
}
