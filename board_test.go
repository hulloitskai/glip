package glip_test

import (
	"github.com/steven-xie/glip"
	"testing"
)

func TestBoard(t *testing.T) {
	b, err := glip.NewBoard()
	if err != nil {
		t.Error("Failed to instantiate Board:", err)
	}

	const in = "Hello, clipboard!"
	if _, err := b.WriteString(in); err != nil {
		t.Error(err)
	}

	out, err := b.ReadString()
	if err != nil {
		t.Error(err)
	}

	if out != in {
		t.Errorf("Expected clipboard paste contents to equal copied string "+
			"(\"%s\"), instead got: \"%s\"", in, out)
	}
}

func TestBoardSafety(t *testing.T) {
	b, err := glip.NewBoard()
	if err != nil {
		t.Error("Failed to instantiate Board:", err)
	}

	b.CopyPortal = nil
	if b.IsWriteable() {
		t.Error("Board should not be writeable with a nil CopyPortal")
	}

	b.PastePortal = nil
	if b.IsReadable() {
		t.Error("Board should not be readable with a nil PastePortal")
	}
}
