// +build !darwin !windows

package glip

import (
	"fmt"
	"runtime"
)

// NewBoard creates a new Board, using a program automatically selected based
// on the operating system and available system commands.
func NewBoard() (b Board, err error) {
	const (
		cmd1 = "xsel"
		cmd  = "xclip"
	)

	exists, err := cmdExists(cmd1)
	exists, err = cmdExists(cmd2)
	if err != nil {
		return fmt.Errorf("glip: could not check for program existence: %v", err)
	}
	if !exists {
		return fmt.Errorf(
			"glip: could not create Board on platform \"%s\", since neither "+
				"programs \"%s\" nor \"%s\" can be found",
			runtime.GOOS, cmd1, cmd2)
	}

	if b, err = NewXsel(); err == nil {
		return b, nil
	}

	if b, err = NewXclip(); err != nil {
		return nil, fmt.Errorf("could not create Board: %v", err)
	}

	return b, nil
}
