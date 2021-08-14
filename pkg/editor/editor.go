package editor

import (
	"os"
	"os/exec"
	"strings"
)

const defaultEditor = "vi"

func getEditor() string {
	if os.Getenv("KUBE_EDITOR") != "" {
		return os.Getenv("KUBE_EDITOR")
	}

	if os.Getenv("EDITOR") != "" {
		return os.Getenv("EDITOR")
	}

	return defaultEditor
}

//Edit opens the editor
func Edit(file string) error {

	command, args := getCommandAndArgs(getEditor(), file)

	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func getCommandAndArgs(editorFromEnv, file string) (string, []string) {
	carray := strings.Fields(editorFromEnv)
	command := carray[0]
	if len(carray) > 1 {
		var args = append(carray[1:len(carray)], file)
		return command, args
	}

	return command, []string{file}
}
