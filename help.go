package main

import (
	"fmt"
	"os"
)

// Print a message to stderr. msg can be of type:
// - string
// - error
func printMsg(msg interface{}) {
	var err error
	switch m := msg.(type) {
	case string:
		// _, err = os.Stderr.WriteString(m + "\n")
		_, err = fmt.Fprintf(os.Stdout, "info:\n%v\n", m)
		if err != nil {
			panic("could not print error message to stderr")
		}
	case error:
		_, err = fmt.Fprintf(os.Stdout, "error:\n%v", m.Error())
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
