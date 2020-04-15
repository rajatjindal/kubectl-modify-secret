package editor

import (
	"os"
	"os/exec"
	"strings"
)

const defaultEditor = "vi"

//Edit opens the editor
func Edit(file string) error {
	editorFromEnv := os.Getenv("EDITOR")
	if editorFromEnv == "" {
		editorFromEnv = defaultEditor
	}

	command, args := getCommandAndArgs(editorFromEnv, file)

	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func getCommandAndArgs(editorFromEnv, file string) (string, []string) {
	carray := strings.Split(editorFromEnv, " ")
	command := carray[0]
	if len(carray) > 1 {
		var args = append(carray[1:len(carray)], file)
		return command, args
	}

	return command, []string{file}
}
