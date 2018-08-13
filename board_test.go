package glip

import "testing"

func TestBoard(t *testing.T) {
	b, err := NewBoard()
	if err != nil {
		t.Error("Failed to instantiate clipboard:", err)
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
