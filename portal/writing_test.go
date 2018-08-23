package portal_test

import (
	"bytes"
	"testing"
)

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
