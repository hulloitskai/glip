// +build !windows

package glip

// Xsel is an API wrapper capable of interfacing with the Xsel Linux program.
//
// Read more about the Xsel program at https://linux.die.net/man/1/xsel.
type Xsel struct {
	// Selection is Xsel's target selection, defaults to XSPrimary.
	//
	// This is where Xsel will save / read clipboard data.
	Selection XSelection

	// Append is the value of Xsel's "--append" flag. If true, Xsel will append
	// data to its target selection (instead of clearing the target selection and
	// replacing its data).
	//
	// Read more about this flag at https://linux.die.net/man/1/xsel.
	Append bool

	*dynPortal
}

// NewXsel creates a new default Xsel instance.
//
// By default, Xsel will target the primary selection (XSPrimary).
func NewXsel() (x *Xsel, err error) {
	return NewXselSelection(XSPrimary)
}

// NewXselSelection creates an Xsel instance targeting a particular X selection.
func NewXselSelection(sel XSelection) (x *Xsel, err error) {
	if err = ensureCmdExists("xsel"); err != nil {
		return nil, err
	}

	x = &Xsel{dynPortal: newDynPortal("xsel"), Selection: sel}
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
