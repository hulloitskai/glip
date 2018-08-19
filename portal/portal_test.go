package portal_test

import (
	"bytes"
	"github.com/steven-xie/glip/portal"
	"testing"
)

func TestPortal_Read(t *testing.T) {
	p := portal.New("echo", TestPhrase)

	data := make([]byte, len(TestPhrase))
	if _, err := p.Read(data); err != nil {
		t.Fatal("Failed to read data from portal:", err)
	}

	outstr := string(data)
	if outstr != TestPhrase {
		t.Errorf("Expected an output of \"%s\", instead got: \"%s\"", TestPhrase,
			outstr)
	}
}

var catp = portal.New("cat")

func TestPortal_ReadFrom(t *testing.T) {
	buf := bytes.NewBufferString(TestPhrase)

	n, err := catp.ReadFrom(buf)
	if err != nil {
		t.Fatal("Could not read from buffer to portal:", err)
	}
	if n == 0 {
		t.Error("Portal unexpectedly read 0 bytes from buffer.")
	}
}

func TestPortal_Write(t *testing.T) {
	n, err := catp.Write([]byte(TestPhrase))
	if err != nil {
		t.Fatal("Could not write to portal:", err)
	}
	if n == 0 {
		t.Error("Portal unexpectedly wrote no data.")
	}
}
