// +build !windows

package glip_test

import (
	"bytes"
	"github.com/steven-xie/glip"
	"testing"
)

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
