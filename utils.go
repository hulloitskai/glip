package glip

import (
	"fmt"
	"os/exec"
)

// cmdExists determines of a particular command can be found in the system
// path.
func cmdExists(name string) (exists bool, err error) {
	if _, err := exec.LookPath(name); err != nil {
		execerr, ok := err.(*exec.Error)
		if !ok || execerr.Err != exec.ErrNotFound {
			return false, fmt.Errorf("glip: error during path lookup: %v", err)
		}
		return false, nil
	}
	return true, nil
}

// ensureCmdExists checks to see if the given command name exists in the path,
// and throws an error if it does not.
func ensureCmdExists(name string) error {
	exists, err := cmdExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("glip: could not find program \"%v\" in system path",
			name)
	}
	return nil
}
