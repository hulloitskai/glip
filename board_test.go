package glip_test

import (
	"bytes"
	"github.com/steven-xie/glip"
	"testing"
)

const TestPhrase = "Hello, clipboard!"

func TestBoard_basic(t *testing.T) {
	b := makeBoard(t)

	if _, err := b.WriteString(TestPhrase); err != nil {
		writeBoardErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		readBoardErr(err, t)
	}

	checkResult(out, t)
}

func TestBoard_Write(t *testing.T) {
	b := makeBoard(t)

	buf := bytes.NewBufferString(TestPhrase)
	if _, err := b.Write(buf.Bytes()); err != nil {
		writeBoardErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		readBoardErr(err, t)
	}

	checkResult(out, t)
}
func TestBoard_ReadFrom(t *testing.T) {
	b := makeBoard(t)

	buf := bytes.NewBufferString(TestPhrase)
	if _, err := b.ReadFrom(buf); err != nil {
		writeBoardErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		readBoardErr(err, t)
	}

	checkResult(out, t)
}

func TestBoard_Read(t *testing.T) {
	b := makeBoard(t)

	if _, err := b.WriteString(TestPhrase); err != nil {
		writeBoardErr(err, t)
	}

	outbuf := make([]byte, len(TestPhrase))
	_, err := b.Read(outbuf)
	if err != nil {
		readBoardErr(err, t)
	}

	checkResult(string(outbuf), t)
}

func TestBoard_WriteTo(t *testing.T) {
	b := makeBoard(t)

	if _, err := b.WriteString(TestPhrase); err != nil {
		writeBoardErr(err, t)
	}
	outbuf := new(bytes.Buffer)

	_, err := b.WriteTo(outbuf)
	if err != nil {
		readBoardErr(err, t)
	}

	checkResult(outbuf.String(), t)
}

func makeBoard(t *testing.T) glip.Board {
	b, err := glip.NewBoard()
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func checkResult(out string, t *testing.T) {
	if out != TestPhrase {
		t.Errorf("Expected output to equal input (\"%s\"), instead got: \"%q\"",
			TestPhrase, out)
	}
}

func readBoardErr(err error, t *testing.T) {
	t.Fatal("Failed to read from clipboard:", err)
}

func writeBoardErr(err error, t *testing.T) {
	t.Fatal("Failed to write to clipboard:", err)
}
