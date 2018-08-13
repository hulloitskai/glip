package glip

import (
	"errors"
	"io"
)

// global is a package-wide Board instance, initialized upon package import. If
// the package encountered an error while initializing the global instance,
// it will instead be nil.
var global *Board

func init() {
	var err error
	if global, err = NewBoard(); err != nil {
		global = nil
	}
}

// checkGlobal returns an error if "global" is nil.
func checkGlobal() error {
	if global == nil {
		return errors.New("failed to initialize package-wide Board " +
			"instance")
	}
	return nil
}

// Read saves data from the system clipboard into the write array, "p".
func Read(p []byte) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return global.Read(p)
}

// ReadString reads the contents of the system clipboard into a string.
func ReadString() (s string, err error) {
	if err = checkGlobal(); err != nil {
		return "", err
	}
	return global.ReadString()
}

// Write records data from "p" into the system clipboard.
func Write(p []byte) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return global.Write(p)
}

// WriteString writes a string into the system clipboard.
func WriteString(s string) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return global.WriteString(s)
}

// WriteTo data from the system clipboard into the given io.Writer.
func WriteTo(w io.Writer) (n int64, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return global.WriteTo(w)
}
