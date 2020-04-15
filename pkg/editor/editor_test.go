package editor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommandAndArgs(t *testing.T) {
	testcases := []struct {
		name            string
		editor          string
		file            string
		expectedCommand string
		expectedArgs    []string
	}{
		{
			name:            "single word editor vim",
			editor:          "vim",
			file:            "some-file.txt",
			expectedCommand: "vim",
			expectedArgs:    []string{"some-file.txt"},
		},
		{
			name:            "single word editor code",
			editor:          "code",
			file:            "some-file.txt",
			expectedCommand: "code",
			expectedArgs:    []string{"some-file.txt"},
		},
		{
			name:            "code with arguments",
			editor:          "code --wait",
			file:            "some-file.txt",
			expectedCommand: "code",
			expectedArgs:    []string{"--wait", "some-file.txt"},
		},
		{
			name:            "code with arguments and inconsistent spacing",
			editor:          "code   --wait",
			file:            "some-file.txt",
			expectedCommand: "code",
			expectedArgs:    []string{"--wait", "some-file.txt"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			command, args := getCommandAndArgs(tc.editor, tc.file)
			assert.Equal(t, tc.expectedCommand, command)
			assert.Equal(t, tc.expectedArgs, args)
		})
	}
}
