package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const version = "1.2"

func main() {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		printErr(errors.New("please set the HOME environment variable"))
	}

	md, err := getMetadata(filepath.Join(homeDir + "/.config/ws"))
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	// Conveniently show the help text if no arguments are passed.
	if len(os.Args) == 1 {
		help()
	} else {
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
		case "create", "c":
			if len(os.Args) != 3 {
				printErr(errors.New("wrong number of arguments"))
			}
			if err := md.insert(os.Args[2]); err != nil {
				printErr(err)
			}
		case "delete", "d":
			if len(os.Args) != 3 {
				printErr(errors.New("wrong number of arguments"))
			}
			if err := md.delete(os.Args[2]); err != nil {
				printErr(err)
			}
		default:
			if i, err := md.getWorkspace(os.Args[1]); err != nil {
				printErr(err)
			} else {
				// The successful return value of this program.
				fmt.Println(md.workspaces[i].Path)
			}
		}
	}
}
