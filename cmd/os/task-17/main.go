package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("current process PID is: %d\n", os.Getpid())
}
