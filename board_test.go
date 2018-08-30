package glip_test

import (
	"bytes"
	"github.com/steven-xie/glip"
	"reflect"
	"testing"
)

const TestPhrase = "Hello, clipboard!"

func TestBoard(t *testing.T) {
	b := makeBoard(t)

	t.Logf("Made a Board, with an underlying type of: %q",
		reflect.TypeOf(b).Elem().Name())

	if _, err := b.WriteString(TestPhrase); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, t)
}

func TestBoard_Write(t *testing.T) {
	b := makeBoard(t)

	buf := bytes.NewBufferString(TestPhrase)
	if _, err := b.Write(buf.Bytes()); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, t)
}
func TestBoard_ReadFrom(t *testing.T) {
	b := makeBoard(t)

	buf := bytes.NewBufferString(TestPhrase)
	if _, err := b.ReadFrom(buf); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, t)
}

func TestBoard_Read(t *testing.T) {
	b := makeBoard(t)

	if _, err := b.WriteString(TestPhrase); err != nil {
		logWriteErr(err, t)
	}

	outbuf := make([]byte, len(TestPhrase))
	_, err := b.Read(outbuf)
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(string(outbuf), t)
}

func TestBoard_WriteTo(t *testing.T) {
	b := makeBoard(t)

	if _, err := b.WriteString(TestPhrase); err != nil {
		logWriteErr(err, t)
	}
	outbuf := new(bytes.Buffer)

	_, err := b.WriteTo(outbuf)
	if err != nil {
		logReadErr(err, t)
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
		t.Errorf("Expected output to equal input (\"%s\"), instead got: %q",
			TestPhrase, out)
	}
}

func logReadErr(err error, t *testing.T) {
	t.Fatal("Failed to read from clipboard:", err)
}

func logWriteErr(err error, t *testing.T) {
	t.Fatal("Failed to write to clipboard:", err)
}
