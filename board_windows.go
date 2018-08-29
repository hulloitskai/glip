// +build windows

package glip

import (
	"fmt"
	"runtime"
)

// NewBoard creates a new Board, using a program automatically selected based
// on the operating system and available system commands.
func NewBoard() (b Board, err error) {
	const (
		cmd1 = "PowerShell"
		cmd2 = "clip"
	)

	// Check for existence of first program, then if first program is not found,
	// check for the second program.
	exists, err := cmdExists(cmd1)
	if !exists && err == nil {
		exists, err = cmdExists(cmd2)
	}

	// If an error occurred during any of the above steps, return it.
	if err != nil {
		return nil, fmt.Errorf("glip: could not check for program existence: %v",
			err)
	}
	if !exists { // none of the programs existed
		return nil, fmt.Errorf(
			"glip: could not create Board on platform \"%s\", since neither "+
				"programs \"%s\" nor \"%s\" can be found",
			runtime.GOOS, cmd1, cmd2)
	}

	if b, err = NewPShellBoard(); err == nil {
		return b, nil
	}

	if b, err = NewWinClip(); err != nil {
		return nil, fmt.Errorf("could not create Board: %v", err)
	}

	return b, nil
}
