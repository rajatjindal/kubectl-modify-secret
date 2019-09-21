package editor

import (
	"os"
	"os/exec"
)

func Edit(file string) error {
	cmd := exec.Command("vim", file)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
