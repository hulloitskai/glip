// +build !windows

package glip

import "io"

// Xsel is an API wrapper capable of interfacing with the Xsel Linux program.
//
// Read more about the Xsel program at https://linux.die.net/man/1/xsel.
type Xsel struct {
	// Append is the value of Xsel's "--append" flag. If true, Xsel will append
	// data to its target selection (instead of clearing the target selection and
	// replacing its data).
	Append bool

	*xPortal
}

const xselPauseLength = 9

// NewXsel creates a new default Xsel instance.
//
// By default, Xsel will target the clipboard selection (XSClipboard).
func NewXsel() (x *Xsel, err error) {
	return NewXselSelection(XSClipboard)
}

// NewXselSelection creates an Xsel instance targeting a particular X selection.
func NewXselSelection(sel XSelection) (x *Xsel, err error) {
	if err = ensureCmdExists("xsel"); err != nil {
		return nil, err
	}

	xp := &xPortal{
		Xopts:       Xopts{Selection: sel},
		dynPortal:   newDynPortal("xsel"),
		pauseLength: xselPauseLength,
	}
	x = &Xsel{xPortal: xp}

	x.GetArgs = x.generateArgs
	return x, nil
}

func (x *Xsel) generateArgs() []string {
	args := []string{"--" + string(x.Selection)}

	if x.Append {
		args = append(args, "--append")
	}
	return args
}

// Keep forces Xsel's primary and secondary selections to persist, using Xsel's
// "--keep" flag.
func (x *Xsel) Keep() error {
	x.AppendArgs("--keep")
	return x.Run()
}

// Exchange exchanges Xsel's primary and secondary selections.
func (x *Xsel) Exchange() error {
	x.AppendArgs("--exchange")
	return x.Run()
}

// Clear clears Xsel's target selection.
func (x *Xsel) Clear() error {
	x.AppendArgs("--clear")
	return x.Run()
}

// Delete requests that Xsel's target selection be deleted, and also requests
// the program in which the selection resides to delete the selected contents.
func (x *Xsel) Delete() error {
	x.AppendArgs("--delete")
	return x.Run()
}

const (
	xselInputFlag  = "--input"
	xselOutputFlag = "--output"
)

// Write writes len(p) bytes into Xsel's target selection.
func (x *Xsel) Write(p []byte) (n int, err error) {
	x.AppendArgs(xselInputFlag)
	return x.xPortal.Write(p)
}

// WriteString writes a string into Xsel's target selection.
func (x *Xsel) WriteString(s string) (n int, err error) {
	x.AppendArgs(xselInputFlag)
	return x.xPortal.WriteString(s)
}

// ReadFrom reads data from an io.Reader into Xsel's target selection.
func (x *Xsel) ReadFrom(r io.Reader) (n int64, err error) {
	x.AppendArgs(xselInputFlag)
	return x.xPortal.ReadFrom(r)
}

// Read reads len(p) bytes from Xsel's target selection into dst.
func (x *Xsel) Read(dst []byte) (n int, err error) {
	x.AppendArgs(xselOutputFlag)
	return x.dynPortal.Read(dst)
}

// WriteTo writes the contents of Xsel's target selection into w.
func (x *Xsel) WriteTo(w io.Writer) (n int64, err error) {
	x.AppendArgs(xselOutputFlag)
	return x.dynPortal.WriteTo(w)
}

// ReadString reads the contents of Xsel's target selection into a string.
func (x *Xsel) ReadString() (s string, err error) {
	x.AppendArgs(xselOutputFlag)
	return x.dynPortal.ReadString()
}
