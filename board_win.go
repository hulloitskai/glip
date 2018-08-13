// +build windows

package glip

import "os/exec"

// NewBoard creates a new Board, if all the necessary system commands ar
// available.
func NewBoard() (b *Board, err error) {
	const (
		copyCmdName  = "copy"
		pasteCmdName = "paste"
	)

	if err = verifyCommands(copyCmdName, pasteCmdName); err != nil {
		return nil, err
	}

	b = MakeBoard(
		exec.Command(copyCmdName),
		exec.Command(pasteCmdName),
	)
	return b, nil
}
