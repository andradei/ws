package main

import (
	"os"
	"os/exec"
)

func main() {
	var osList = []string{"darwin", "linux"}
	var archList = []string{"amd64", "386"}

	for _, goos := range osList {

		for _, goarch := range archList {
			var cmd = exec.Command("go", "build", "-o", "bin/ws_"+goos+"_"+goarch)
			cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)

			if err := cmd.Run(); err != nil {
				panic("couldn't run command: " + err.Error())
			}
		}

	}
}
