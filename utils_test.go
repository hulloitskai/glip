package glip

import "testing"

func TestCmdExists(t *testing.T) {
	name := "ls"
	exists, err := cmdExists(name)
	if err != nil {
		t.Errorf("Error while checking if program \"%s\" exists: %v", name, err)
	}
	if !exists {
		t.Errorf("Expected command \"%s\" to exist.", name)
	}

	name = "ls11111111111111111"
	exists, err = cmdExists(name)
	if err != nil {
		t.Errorf("Error while checking if program \"%s\" exists: %v", name, err)
	}
	if exists {
		t.Errorf("Did not expect command \"%s\" to exist.", name)
	}
}
