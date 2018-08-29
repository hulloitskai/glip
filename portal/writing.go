package portal

import (
	"fmt"
	"io"
)

// Write writes len(data) bytes into the Portal.
func (p *Portal) Write(data []byte) (n int, err error) {
	defer p.Reload()

	// Open a pipe to program Stdin.
	in, err := p.stdinPipe()
	if err != nil {
		return 0, err
	}

	// Start program; begin writing to it's Stdin from data.
	if err = p.start(); err != nil {
		return 0, err
	}
	if n, err = in.Write(data); err != nil {
		return n, fmt.Errorf("portal: error while writing to Stdin: %v", err)
	}

	// Close Stdin to signal to the program that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	// Wait for the program to exit.
	return n, p.wait()
}

// ReadFrom reads data from an io.Reader into the Portal.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	defer p.Reload()

	// Open a pipe to Stdin.
	in, err := p.stdinPipe()
	if err != nil {
		return 0, err
	}

	// Start the program; read from the provided io.Reader to program Stdin.
	if err = p.start(); err != nil {
		return 0, err
	}
	if n, err = io.Copy(in, r); err != nil {
		return n, fmt.Errorf("portal: error while copying to Stdin: %v", err)
	}

	// Close program Stdin to signal that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	// Wait for program to exit.
	return n, p.wait()
}
