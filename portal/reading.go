package portal

import (
	"fmt"
	"io"
)

// Read reads len(dst) of data from the Portal into dst.
func (p *Portal) Read(dst []byte) (n int, err error) {
	defer p.Reload()

	// Open an pipe to Stdout.
	out, err := p.stdoutPipe()
	if err != nil {
		return 0, err
	}

	// Start p.Cmd; read data to destination.
	if err = p.start(); err != nil {
		return 0, err
	}
	if n, err = out.Read(dst); err != nil {
		return n, fmt.Errorf("portal: error while reading from Stdout: %v", err)
	}

	// Wait for p.Cmd to complete.
	return n, p.wait()
}

// WriteTo writes data from the Portal into an io.Writer.
func (p *Portal) WriteTo(w io.Writer) (n int64, err error) {
	defer p.Reload()

	// Open a pipe to Stdout.
	out, err := p.stdoutPipe()
	if err != nil {
		return 0, err
	}

	// Start p.Cmd; copy data from program Stdout to the provided io.Writer.
	if err = p.start(); err != nil {
		return 0, err
	}
	if n, err = io.Copy(w, out); err != nil {
		return n, fmt.Errorf("portal: error while copying from Stdout: %v", err)
	}

	// Wait for p.Cmd to complete.
	return n, p.wait()
}
