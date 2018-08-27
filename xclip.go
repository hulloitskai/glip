// +build !windows

package glip

import (
	"io"
	"strconv"
)

// Xclip is an API wrapper capable of interfacing with the Xclip Linux program.
//
// Read more about this flag at https://linux.die.net/man/1/xclip.
type Xclip struct {
	// Selection is Xclip's target selection.
	Selection XSelection

	// Filter represents the Xclip's "-filter" flag.
	//
	// It causes Xclip to print the text piped to standard in back to standard
	// out, unmodified.
	Filter bool

	// Loops represents Xclip's "-loops" flag.
	//
	// If this value is non-zero, Xclip will wait this many X selection requests
	// (pastes into X applications) before exiting.
	Loops uint

	// Display represents Xclip's "-display" flag.
	//
	// If not set (the zero value), it defaults to the value in the "$DISPLAY"
	// environment variable.
	Display string

	*dynPortal
}

// NewXclip creates a new default Xclip instance.
//
// By default, none of Xclip's flags are enabled, and Xclip will use the
// XSPrimary selection.
func NewXclip() (x *Xclip, err error) {
	return NewXclipSelection(XSPrimary)
}

// NewXclipSelection creates an Xclip instance targeting a particular X
// selection.
func NewXclipSelection(sel XSelection) (x *Xclip, err error) {
	if err = ensureCmdExists("xclip"); err != nil {
		return nil, err
	}

	x = &Xclip{dynPortal: newDynPortal("xclip"), Selection: sel}
	x.GetArgs = x.generateArgs
	return x, nil
}

func (x *Xclip) generateArgs() []string {
	args := []string{"-sel", string(x.Selection)}

	if x.Filter {
		args = append(args, "-filter")
	}
	if x.Loops != 0 {
		args = append(args, "-loops", strconv.FormatUint(uint64(x.Loops), 10))
	}
	if x.Display != "" {
		args = append(args, "-display", x.Display)
	}

	return args
}

const xclipOutFlag = "-out"

// Write writes data to an X selection.
func (x *Xclip) Write(data []byte) (n int, err error) {
	x.AppendArgs(xclipOutFlag)
	return x.dynPortal.Write(data)
}

// ReadFrom reads data from an io.Reader into an X selection.
func (x *Xclip) ReadFrom(r io.Reader) (n int64, err error) {
	x.AppendArgs(xclipOutFlag)
	return x.dynPortal.ReadFrom(r)
}
