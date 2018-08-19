// +build windows

package glip

import "os/exec"

// NewBoard creates a new Board, if all the necessary system commands ar
// available.
func NewBoard() (b *Board, err error) {
	const copyCmdName = "clip"
	const pasteCmdName = "paste"

	if err = verifyCommand(copyCmdName); err != nil {
		return nil, err
	}

	var (
		copyCmd  = exec.Command(copyCmdName)
		pasteCmd *exec.Cmd
	)
	if err = verifyCommand(pasteCmdName); err != nil {
		pasteCmd = exec.Command(pasteCmdName)
	}

	b = MakeBoard(copyCmd, pasteCmd)
	return b, nil
}
