package portal

import (
	"fmt"
	"io"
	"log"
)

// Read reads len(dst) of data from the Portal into dst.
func (p *Portal) Read(dst []byte) (n int, err error) {
	defer p.Reload()

	log.Print("portal.Read: begun")

	// Open an pipe to Stdout.
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdoutPipe: %v", err)
	}

	log.Print("portal.Read: got Stdout")

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	log.Print("portal.Read: started program")

	// Perform read operation.
	if n, err = out.Read(dst); err != nil {
		return 0, fmt.Errorf("portal: error while reading from Stdout: %v", err)
	}

	log.Print("portal.Read: read from Stdout")

	// Wait for Cmd to complete.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	log.Print("portal.Read: program exited")

	return n, nil
}

// WriteTo writes data from the Portal into an io.Writer.
func (p *Portal) WriteTo(w io.Writer) (n int64, err error) {
	defer p.Reload()

	log.Print("portal.WriteTo: begun")

	// Open a pipe to Stdout.
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdoutPipe: %v", err)
	}

	log.Print("portal.WriteTo: got Stdout")

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	log.Print("portal.WriteTo: started program")

	// Copy data from Stdout into w.
	if n, err = io.Copy(w, out); err != nil {
		return 0, fmt.Errorf("portal: error while copying from Stdout: %v", err)
	}

	log.Print("portal.WriteTo: copied from Stdout")

	// Wait for Cmd to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	log.Print("portal.WriteTo: program exited")

	return n, nil
}
