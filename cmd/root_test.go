package cmd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_execute(t *testing.T) {
	tcs := map[string]struct {
		command     string
		expectError bool
	}{
		"valid command": {
			command:     "help",
			expectError: false,
		},
		"not supported command fallbacks to git": {
			command:     "status -s",
			expectError: false,
		},
	}
	for name, tc := range tcs {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			root := newRootCmd()
			args := strings.Split(tc.command, " ")
			root.SetArgs(args)
			err := execute(root, args...)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_fallbackToGit(t *testing.T) {
	tcs := map[string]struct {
		command     string
		expectError bool
	}{
		"invalid git command return error": {
			command:     "xpto",
			expectError: true,
		},
		"valid git command with invalid flag return error": {
			command:     "status -buaua",
			expectError: true,
		},
		"valid git command": {
			command:     "status",
			expectError: false,
		},
		"valid git command with valid flag return error": {
			command:     "status -s",
			expectError: false,
		},
	}

	for name, tc := range tcs {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := fallbackToGit(strings.Split(tc.command, " ")...)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
