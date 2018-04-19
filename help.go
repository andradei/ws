package main

import (
	"fmt"
	"os"

	color "github.com/fatih/color-1.5.0"
)

// Print an error to stderr and exits the program. May panic.
func printErr(msg error) {
	_, err := fmt.Fprint(os.Stderr, color.RedString("ws: %v\n", msg))
	if err != nil {
		panic("could not print error message to stderr")
	}
	os.Exit(1)
}

// Display tool usage information.
func help() {
	// TODO: Turn this into a template
	// TODO: Register functions of color packages to print with color
	cmdFmt := "\n  %v | %v %v\n    %v\n"
	ws := color.YellowString("ws")
	workspaceName := color.YellowString("workspace name")
	title := fmt.Sprintf(
		"\n%v - %v (v%v)\n",
		ws,
		color.WhiteString("Directory alias manager"),
		version,
	)
	usage := fmt.Sprintf(
		"\n%v %v <%v [%v]> | <%v>\n",
		color.GreenString("Usage:"), ws, color.YellowString("command"), workspaceName, workspaceName,
	)
	create := fmt.Sprintf(cmdFmt, color.YellowString("create"), color.YellowString("c"), workspaceName, "Create a workspace with given name")
	delete := fmt.Sprintf(cmdFmt, color.YellowString("delete"), color.YellowString("d"), workspaceName, "Delete an existing workspace by name")
	help := fmt.Sprintf(cmdFmt, color.YellowString("help"), color.YellowString("h"), workspaceName, "Display this help message")
	list := fmt.Sprintf(cmdFmt, color.YellowString("list"), color.YellowString("l"), workspaceName, "List existing workspaces")

	fmt.Println(title, usage, create, delete, help, list)

	exampleFmt := "\n  %v:\n    %v\n"
	examples := color.GreenString("\nExamples:\n")
	example1 := fmt.Sprintf(exampleFmt, color.YellowString("Create workspace"), "ws create project1")
	example2 := fmt.Sprintf(exampleFmt, color.YellowString("Go to workspace"), "cd $(ws project1)")
	example3 := fmt.Sprintf(exampleFmt, color.YellowString("Delete workspace"), "ws delete project1")

	fmt.Println(examples, example1, example2, example3)
}
