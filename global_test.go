package glip_test

import (
	"bytes"
	"fmt"
	"github.com/steven-xie/glip"
	"testing"
)

func Example() {
	// Write a string into the clipboard.
	glip.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// the input string.
	out, err := glip.ReadString()
	if err != nil {
		fmt.Printf("Encountered an error while reading from clipboard: %v", err)
	}

	fmt.Println(out)
	// Output: example string
}

func TestGlobal(t *testing.T) {
	const instr = "A B C D E F G"
	inbuf := bytes.NewBufferString(instr)
	if _, err := glip.ReadFrom(inbuf); err != nil {
		t.Fatal("Failed to write to clipboard:", err)
	}

	outbuf := new(bytes.Buffer)
	if _, err := glip.WriteTo(outbuf); err != nil {
		t.Fatal("Failed to read from clipboard:", err)
	}

	if outstr := outbuf.String(); outstr != instr {
		t.Fatalf("Expected output of \"%s\", instead got: \"%s\"", instr, outstr)
	}
}
