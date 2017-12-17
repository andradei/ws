package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	md, err := getMetadata()
	if err != nil {
		errorMsg(err)
		os.Exit(1)
	}
	// Check the quantity of arguments are the allowed quantity.
	l := len(os.Args)
	switch l {
	// Show help if no command is given
	case 1:
		// TODO: Got to a default (or the last) workspace?
		help()
	case 2:
		// The argument is either a known command or a workspace name.
		switch os.Args[1] {
		case "-help", "-h":
			help()
		case "-list", "-ls", "-l":
			ls, err := md.list()
			if err != nil {
				errorMsg(fmt.Errorf("list: %v", err))
			}
			fmt.Println(ls)
		default:
			ws := os.Args[1]
			if strings.HasPrefix(ws, "-") {
				errorMsg(fmt.Errorf("command %s not found", ws))
			} else {
				if i, err := md.getWorkspace(ws); err != nil {
					errorMsg(err)
				} else {
					// The successful return value of this program is the workspace's path.
					fmt.Print(md.workspaces[i].Path)
				}
			}
		}
	case 3:
		cmd := os.Args[1]
		ws := os.Args[2]

		switch cmd {
		case "-insert", "-i":
			// Get current working directory
			pwd, err := os.Getwd()
			if err != nil {
				errorMsg(fmt.Errorf("unable to retrieve working directory: %v", err))
			}
			if err := md.insert(ws, pwd); err != nil {
				errorMsg(err)
			}
		case "-delete", "-d":
			if err := md.delete(ws); err != nil {
				errorMsg(err)
			}
		default:
			errorMsg(errors.New("please provide a valid command (See -help or -h)"))
		}
	}
}
