package glip

import (
	"fmt"
	"io"
)

var (
	// B is a package-wide Board instance, initialized upon package import.
	//
	// If an error was encountered while initializing B, B will instead be nil.
	B Board

	// BErr represents the error that was encountered when attempting to
	// initialize B, the package-wide Board instance.
	//
	// It is nil if the Board instance was created successfully.
	BErr error
)

func init() {
	if B, BErr = NewBoard(); BErr != nil {
		B = nil
	}
}

// checkGlobal returns an error if B, the package-wide Board, is nil.
func checkGlobal() error {
	if BErr != nil {
		return fmt.Errorf("glip: failed to initialize package-wide Board "+
			"instance: %v", BErr)
	}
	return nil
}

////////////////////////////
// Reading from B
////////////////////////////

// Read reads len(dst) bytes from the system clipboard into dst.
func Read(dst []byte) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return B.Read(dst)
}

// ReadString reads the contents of the system clipboard into a string.
func ReadString() (s string, err error) {
	if err = checkGlobal(); err != nil {
		return "", err
	}
	return B.ReadString()
}

// WriteTo writes data from the system clipboard into an io.Writer.
func WriteTo(w io.Writer) (n int64, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return B.WriteTo(w)
}

////////////////////////////
// Writing to B
////////////////////////////

// Write writes len(dst) bytes into the system clipboard.
func Write(dst []byte) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return B.Write(dst)
}

// WriteString writes a string into the system clipboard.
func WriteString(s string) (n int, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return B.WriteString(s)
}

// ReadFrom reads data from an io.Reader into the system clipboard.
func ReadFrom(r io.Reader) (n int64, err error) {
	if err = checkGlobal(); err != nil {
		return 0, err
	}
	return B.ReadFrom(r)
}
