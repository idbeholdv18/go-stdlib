package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(checkIsFileExists("./test.txt"))
	fmt.Println(checkIsFileExists("./test.go"))
}

func checkIsFileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if info.IsDir() {
		fmt.Printf("%s is dir\n", info.Name())
		return false, nil
	}
	return true, nil
}
