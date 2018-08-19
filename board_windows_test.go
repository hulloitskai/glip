// +build windows

package glip_test

import (
	"github.com/steven-xie/glip"
	"testing"
)

const TestPhrase = "Hello, clipboard!"

func TestBoard(t *testing.T) {
	b, err := glip.NewBoard()
	if err != nil {
		t.Fatal("Failed to instantiate Board:", err)
	}

	n, err := b.WriteString(TestPhrase)
	if err != nil {
		t.Errorf("Failed to write string (\"%s\") to clipboard: %v", TestPhrase,
			err)
	}
	if n == 0 {
		t.Errorf("Unexpectedly got non-zero bytes written (%d bytes)", n)
	}
}
