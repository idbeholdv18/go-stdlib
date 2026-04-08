package main

import (
	"os/exec"
)

func main() {
	copyCmd := exec.Command("pbcopy")

	stdin, err := copyCmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	if err := copyCmd.Start(); err != nil {
		panic(err)
	}

	cmd := exec.Command("ps", "-elf")
	cmd.Stdout = stdin

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	stdin.Close()

	cmd.Wait()
}
