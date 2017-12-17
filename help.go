package main

import (
	"os"
)

// Print a message to stderr.
func errorMsg(msg error) {
	_, err := os.Stderr.WriteString(msg.Error() + "\n")
	if err != nil {
		panic("could not print error message to stderr")
	}
}

// Display tool usage information.
// TODO
func help() {
	panic("unimplemented")
}
