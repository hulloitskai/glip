package portal

import (
	"fmt"
	"io"
	"log"
)

// Write writes len(data) bytes into the Portal.
func (p *Portal) Write(data []byte) (n int, err error) {
	defer p.Reload()

	log.Print("portal.Write: begun")

	// Open a pipe to Stdin.
	in, err := p.StdinPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdinPipe: %v", err)
	}

	log.Print("portal.Write: opened Stdin")

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	log.Print("portal.Write: started program")

	// Perform write operation.
	if n, err = in.Write(data); err != nil {
		return 0, fmt.Errorf("portal: error while writing to Stdin: %v", err)
	}

	log.Print("portal.Write: wrote to Stdin")

	// Close Stdin to signal to the program that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	log.Print("portal.Write: close Stdin")

	// Wait for the program to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	log.Print("portal.Write: program exited")

	return n, nil
}

// ReadFrom writes data from an io.Reader into the Portal.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	defer p.Reload()

	log.Print("portal.ReadFrom: begun")

	// Open a pipe to program stdin.
	in, err := p.StdinPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdinPipe: %v", err)
	}

	log.Print("portal.ReadFrom: got Stdin")

	// Start the program.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	log.Print("portal.ReadFrom: started program")

	// Copy data from r into Stdin.
	if n, err = io.Copy(in, r); err != nil {
		return 0, fmt.Errorf("portal: error while copying to Stdin: %v", err)
	}

	log.Print("portal.ReadFrom: copied to Stdin")

	// Close program stdin to signal that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	log.Print("portal.ReadFrom: closed Stdin")

	// Wait for program to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}

	log.Print("portal.ReadFrom: program exited")

	return n, nil
}
