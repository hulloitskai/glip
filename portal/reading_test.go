package portal_test

import (
	"testing"
)

func TestPortal_Read(t *testing.T) {
	data := make([]byte, len(TestPhrase))
	if _, err := echop.Read(data); err != nil {
		t.Fatal("Failed to read data from portal:", err)
	}

	outstr := string(data)
	if outstr != TestPhrase {
		t.Errorf("Expected an output of \"%s\", instead got: \"%s\"", TestPhrase,
			outstr)
	}
}
