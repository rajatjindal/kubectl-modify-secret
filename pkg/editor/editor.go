package editor

import (
	"os"
	"os/exec"
)

const defaultEditor = "vi"

//Edit opens the editor
func Edit(file string) error {
	editorFromEnv := os.Getenv("EDITOR")
	if editorFromEnv == "" {
		editorFromEnv = defaultEditor
	}

	cmd := exec.Command(editorFromEnv, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
