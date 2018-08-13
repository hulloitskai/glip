package glip

import "testing"

func TestVerifyCommand(t *testing.T) {
	if verifyCommand("ls") != nil {
		t.Error("Expected command \"ls\" to exist.")
	}

	if verifyCommand("ls11111") == nil {
		t.Error("Did not expect command \"ls11111\" to exist.")
	}
}
