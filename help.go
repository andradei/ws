package main

import (
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
)

// Print an error to stderr and exits the program. May panic.
func printErr(msg error) {
	_, err := fmt.Fprint(os.Stderr, aurora.Sprintf(aurora.Red("ws: %v\n"), msg))
	if err != nil {
		panic("could not print error message to stderr")
	}
	os.Exit(1)
}

// Display tool usage information.
func help() {
	// TODO: Turn this into a template
	// TODO: Register functions of color packages to print with color
	cmdFmt := "\n  %v | %v <%v>\n    %v\n"
	ws := fmt.Sprint(aurora.Brown("ws"))
	workspaceName := fmt.Sprint(aurora.Brown("workspace name"))
	title := fmt.Sprintf(
		"\n%v - %v (v%v)\n",
		ws,
		aurora.Gray("Directory alias manager"),
		version,
	)
	usage := fmt.Sprintf(
		"\n%v %v <%v [%v]> | <%v>\n",
		aurora.Green("Usage:"), ws, aurora.Brown("command"), workspaceName, workspaceName,
	)
	create := fmt.Sprintf(cmdFmt, aurora.Brown("create"), aurora.Brown("c"),
		workspaceName, "Create a workspace with given name")
	delete := fmt.Sprintf(cmdFmt, aurora.Brown("delete"), aurora.Brown("d"),
		workspaceName, "Delete an existing workspace by name")
	help := fmt.Sprintf(cmdFmt, aurora.Brown("help"), aurora.Brown("h"),
		workspaceName, "Display this help message")
	list := fmt.Sprintf(cmdFmt, aurora.Brown("list"), aurora.Brown("l"),
		workspaceName, "List existing workspaces")

	fmt.Println(title, usage, create, delete, help, list)

	exampleFmt := "\n  %v:\n    %v\n"
	examples := aurora.Green("\nExamples:\n")
	example1 := fmt.Sprintf(exampleFmt, aurora.Brown("Create workspace"), "ws create project1")
	example2 := fmt.Sprintf(exampleFmt, aurora.Brown("Go to workspace"), "cd $(ws project1)")
	example3 := fmt.Sprintf(exampleFmt, aurora.Brown("Delete workspace"), "ws delete project1")

	fmt.Println(examples, example1, example2, example3)
}
