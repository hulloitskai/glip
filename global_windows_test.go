package glip_test

import (
	"github.com/steven-xie/glip"
	"testing"
)

func TestGlobal(t *testing.T) {
	n, err := glip.WriteString(TestPhrase)
	if err != nil {
		t.Error("Failed to write string to clipboard:", err)
	}
	if n == 0 {
		t.Errorf("Unexpectedly got a non-zero bytes written (%d bytes)", n)
	}
}
