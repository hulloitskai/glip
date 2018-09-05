package portal

import (
	"fmt"
	"io"
)

// Read reads len(dst) of data from the Portal into dst.
func (p *Portal) Read(dst []byte) (n int, err error) {
	defer p.Reload()

	// Open an pipe to Stdout.
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdoutPipe: %v", err)
	}

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	// Perform read operation.
	if n, err = out.Read(dst); err != nil {
		return 0, fmt.Errorf("portal: error while reading from Stdout: %v", err)
	}

	// Wait for Cmd to complete.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	return n, nil
}

// WriteTo writes data from the Portal into an io.Writer.
func (p *Portal) WriteTo(w io.Writer) (n int64, err error) {
	defer p.Reload()

	// Open a pipe to Stdout.
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdoutPipe: %v", err)
	}

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	// Copy data from Stdout into w.
	if n, err = io.Copy(w, out); err != nil {
		return 0, fmt.Errorf("portal: error while copying from Stdout: %v", err)
	}

	// Wait for Cmd to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	return n, nil
}
