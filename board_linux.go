// +build linux openbsd netbsd solaris

package glip

import (
	"errors"
	"fmt"
	"os/exec"
)

// chooseLinuxCommand checks for the existence of "xclip" and "xsel", returning
// the name of the first one it finds. It returns an error of none are found.
func chooseLinuxCommand() (name string, err error) {
	if _, err := exec.LookPath("xclip"); err == nil {
		return "xclip", nil
	}

	if _, err := exec.LookPath("xsel"); err == nil {
		return "xsel", nil
	}

	return name, errors.New("failed to locate one of \"xclip\" or \"xsel\" in " +
		"the system path")
}

// LinuxCmd refers to a Linux-specific clipboard command.
type LinuxCmd = string

const (
	// Auto allows glip to choose the Linux clipboard command (one of "xclip",
	// "xsel"), whichever is available on the system.
	Auto LinuxCmd = "auto"
	// Xclip makes glip use the "xclip" Linux clipboard utility.
	Xclip LinuxCmd = "xclip"
	// Xsel makes glip use the "xsel" Linux utility.
	Xsel LinuxCmd = "xsel"
	// LinuxDefaultClipboard is the default Linux clipboard. This can be used
	// instead of the "clipboard" argument for NewLinuxBoard.
	LinuxDefaultClipboard = ""
)

// NewLinuxBoard creates a customized Board using a custom command name, and
// a custom target clipboard (as it is possible to have multiple clipboards
// on Linux).
//
// The "args" argument is a set of extra arguments to be passed to the clipboard
// commands.
func NewLinuxBoard(lc LinuxCmd, clipboard string, args ...string) (
	b *Board, err error,
) {
	var (
		cmdName string
		cmdArgs []string
	)

	// Determine cmdName...
	switch lc {
	case Auto:
		if cmdName, err = chooseLinuxCommand(); err != nil {
			return nil, err
		}
	case Xclip, Xsel:
		cmdName = lc
	default:
		err = fmt.Errorf("unable to create clipboard from unknown command: %s",
			cmdName)
		return nil, err
	}

	// Add clipboard string to args...
	if clipboard != LinuxDefaultClipboard {
		switch cmdName {
		case "xclip":
			cmdArgs = append([]string{"-sel", clipboard}, args...)
		case "xsel":
			cmdArgs = append([]string{"--" + clipboard}, args...)
		default:
			err = fmt.Errorf("unable to create clipboard from unknown command: %s",
				cmdName)
		}
	}

	var (
		copyCmd  *exec.Cmd
		pasteCmd *exec.Cmd
	)

	switch cmdName {
	case "xclip":
		copyCmd = exec.Command(cmdName, cmdArgs...)
		pasteCmd = exec.Command(cmdName, append([]string{"-o"}, cmdArgs...)...)
	case "xsel":
		copyCmd = exec.Command(cmdName, cmdArgs...)
		pasteCmd = copyCmd
	default:
		err = fmt.Errorf("unable to create clipboard from unknown command: %s",
			cmdName)
		return nil, err
	}

	return MakeBoard(copyCmd, pasteCmd), nil
}

// NewBoard creates a new Board, if all the necessary system commands ar
// available.
func NewBoard() (b *Board, err error) {
	return NewLinuxBoard(Auto, LinuxDefaultClipboard)
}
