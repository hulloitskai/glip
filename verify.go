package glip

import (
	"fmt"
	"os/exec"
)

// verifyCommand returns an error if the given command name does not exist
// in system path.
func verifyCommand(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		return fmt.Errorf("could not find command with name: %s", name)
	}
	return nil
}

// verifyCommands verifies a set of commands using "verifyCommand". Errors out
// if any of the provided command names are not found in the system path.
func verifyCommands(names ...string) error {
	for _, name := range names {
		if err := verifyCommand(name); err != nil {
			return err
		}
	}
	return nil
}
