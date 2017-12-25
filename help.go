package main

import (
	"os"
)

// Print a message to stderr. msg can be of type:
// - string
// - error
func printMsg(msg interface{}) {
	var err error
	switch m := msg.(type) {
	case string:
		_, err = os.Stderr.WriteString(m + "\n")
		if err != nil {
			panic("could not print error message to stderr")
		}
	case error:
		_, err = os.Stderr.WriteString(m.Error() + "\n")
		if err != nil {
			panic("could not print error message to stderr")
		}
	}
}

// Display tool usage information.
// TODO
func help() {
	panic("unimplemented")
}
