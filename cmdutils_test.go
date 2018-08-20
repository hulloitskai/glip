package glip

import "testing"

func TestCmdBuilder_Verify(t *testing.T) {
	ls := newCmdBuilder("ls")

	found, err := ls.verify()
	if err != nil {
		t.Errorf("Error while verifying program (%s): %v", ls.name, err)
	}
	if !found {
		t.Errorf("Expected command \"%s\" to exist.", ls.name)
	}

	dummy := newCmdBuilder("ls1111111")
	if found, err = dummy.verify(); err != nil {
		t.Errorf("Error while verifying program (%s): %v", dummy.name, err)
	}
	if found {
		t.Errorf("Did not expect command \"%s\" to exist.", dummy.name)
	}
}
