package editor

import (
	"errors"
	"os"
	"os/exec"
)

//Edit opens the editor
func Edit(file string) error {
	editorFromEnv := os.Getenv("EDITOR")
	if editorFromEnv == "" {
		return errors.New("ENV variable $EDITOR not set")
	}

	cmd := exec.Command(editorFromEnv, file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
