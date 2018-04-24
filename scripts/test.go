package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var cmd *exec.Cmd

	fmt.Println("-- Running tests and generating coverage profile...")

	cmd = exec.Command("go", "test", "-coverprofile", "coverage.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		exit("error running test command:", err)
	}

	fmt.Println("-- Generating HTML coverage report...")

	// Reuse the cmd variable for the next command.
	cmd = exec.Command("go", "tool", "cover", "-html", "coverage.out", "-o", "coverage.html")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		exit("error running coverage command:", err)
	}

	fmt.Println("-- All tasks done.")
}

func exit(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}
