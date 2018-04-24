package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const version = "1.2"

func main() {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		printErr(errors.New("please set the HOME environment variable"))
	}

	md, err := getMetadata(filepath.Join(homeDir + ".config/ws"))
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	// Check the quantity of arguments are the allowed quantity.
	l := len(os.Args)
	switch l {
	case 1:
		// Show help if no command is given
		help()
	case 2:
		// The argument is either a known command or a workspace name.
		switch os.Args[1] {
		case "help", "h":
			help()
		case "list", "l":
			ls, err := md.list()
			if err != nil {
				printErr(fmt.Errorf("list: %v", err))
			}
			fmt.Print(ls)
		case "version", "v":
			fmt.Printf("ws v%s\n", version)
		default:
			ws := os.Args[1]
			if strings.HasPrefix(ws, "-") {
				cmds := []string{"delete", "d", "create", "c"}
				for _, cmd := range cmds {
					if ws == cmd {
						printErr(fmt.Errorf("command %s requires an aditional argument, see \"ws help\"", cmd))
					}
				}
				printErr(fmt.Errorf("command %s not found", ws))
			} else {
				if i, err := md.getWorkspace(ws); err != nil {
					printErr(err)
				} else {
					// The successful return value of this program is the workspace's path.
					fmt.Println(md.workspaces[i].Path)
				}
			}
		}
	case 3:
		cmd := os.Args[1]
		ws := os.Args[2]

		switch cmd {
		case "create", "c":
			// Get current working directory
			pwd, err := os.Getwd()
			if err != nil {
				printErr(fmt.Errorf("unable to retrieve working directory: %v", err))
			}
			if err := md.insert(ws, pwd); err != nil {
				printErr(err)
			}
		case "delete", "d":
			if err := md.delete(ws); err != nil {
				printErr(err)
			}
		default:
			printErr(errors.New("please provide a valid command (See \"ws help\")"))
		}
	}
}
