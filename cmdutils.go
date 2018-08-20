package glip

import (
	"fmt"
	"os/exec"
)

// cmdBuilder is a struct used to construct exec.Cmds, but which uses less
// memory than a raw exec.Cmd.
type cmdBuilder struct {
	name string
	args []string
}

// newCmdBuilder makes a new cmdBuilder.
func newCmdBuilder(name string, args ...string) cmdBuilder {
	return cmdBuilder{name, args}
}

// build creates an exec.Cmd out of a cmdBuilder's fields.
func (cb *cmdBuilder) build() *exec.Cmd {
	return exec.Command(cb.name, cb.args...)
}

// verify checks if a cmdBuilder can be found in the system path.
func (cb *cmdBuilder) verify() (found bool, err error) {
	if _, err := exec.LookPath(cb.name); err != nil {
		execerr, ok := err.(*exec.Error)
		if !ok || execerr.Err != exec.ErrNotFound {
			return false, fmt.Errorf("glip: error during path lookup: %v", err)
		}
		return false, nil
	}
	return true, nil
}

// autoselectCB returns the index of the first cmdBuilder that can be found in
// the system path.
//
// It returns -1 if no such cmdBuilder is valid.
func autoselectCB(builders []cmdBuilder) (index int, err error) {
	for i, cb := range builders {
		found, err := cb.verify()
		if err != nil {
			return -1, err
		}
		if found {
			return i, nil
		}
	}
	return -1, nil
}
