package glip_test

import (
	"bytes"
	"github.com/stevenxie/glip"
	"reflect"
	"testing"
)

func TestBoard(t *testing.T) {
	const instr = "TestBoard"
	b := makeBoard(t)

	t.Logf("Made a Board, with an underlying type of: %q",
		reflect.TypeOf(b).Elem().Name())

	if _, err := b.WriteString(instr); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, instr, t)
}

func TestBoard_multilength(t *testing.T) {
	b := makeBoard(t)

	for i := 100; i < 1000; i += 100 {
		in := genstr(i)
		t.Logf("Testing clipboard IO with string: \"%v\"", in)

		if _, err := b.WriteString(in); err != nil {
			t.Errorf("Failed to write string of length %d: %v", i, err)
		}

		out, err := b.ReadString()
		if err != nil {
			t.Errorf("Failed to read string of length %d: %v", i, err)
		} else {
			checkResult(out, in, t)
		}
	}
}

func TestBoard_Write(t *testing.T) {
	const instr = "TestBoard_Write"
	b := makeBoard(t)

	buf := bytes.NewBufferString(instr)
	if _, err := b.Write(buf.Bytes()); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, instr, t)
}
func TestBoard_ReadFrom(t *testing.T) {
	const instr = "TestBoard_ReadFrom"
	b := makeBoard(t)

	buf := bytes.NewBufferString(instr)
	if _, err := b.ReadFrom(buf); err != nil {
		logWriteErr(err, t)
	}

	out, err := b.ReadString()
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(out, instr, t)
}

func TestBoard_Read(t *testing.T) {
	const instr = "TestBoard_Read"
	b := makeBoard(t)

	n, err := b.WriteString(instr)
	if err != nil {
		logWriteErr(err, t)
	}

	outbuf := make([]byte, n)
	if _, err = b.Read(outbuf); err != nil {
		logReadErr(err, t)
	}

	checkResult(string(outbuf), instr, t)
}

func TestBoard_WriteTo(t *testing.T) {
	const instr = "TestBoard_WriteTo"
	b := makeBoard(t)

	if _, err := b.WriteString(instr); err != nil {
		logWriteErr(err, t)
	}
	outbuf := new(bytes.Buffer)

	_, err := b.WriteTo(outbuf)
	if err != nil {
		logReadErr(err, t)
	}

	checkResult(outbuf.String(), instr, t)
}

func makeBoard(t *testing.T) glip.Board {
	b, err := glip.NewBoard()
	if err != nil {
		t.Fatal(err)
	}

	if xb, ok := b.(glip.XBoard); ok {
		xb.Opts().SafeWrites = true
	}

	return b
}

func checkResult(out string, expect string, t *testing.T) {
	if out != expect {
		t.Errorf("Expected output to equal input (\"%s\"), instead got: \"%s\"",
			expect, out)
	}
}

func logReadErr(err error, t *testing.T) {
	t.Fatal("Failed to read from clipboard:", err)
}

func logWriteErr(err error, t *testing.T) {
	t.Fatal("Failed to write to clipboard:", err)
}

// genstr generates a string of the specified length.
func genstr(length int) string {
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		data[i] = 'E'
	}
	return string(data)
}
