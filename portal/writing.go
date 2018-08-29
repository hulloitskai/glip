package portal

import (
	"fmt"
	"io"
)

// Write writes len(data) bytes into the Portal.
func (p *Portal) Write(data []byte) (n int, err error) {
	defer p.Reload()

	// Open a pipe to Stdin.
	in, err := p.StdinPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdinPipe: %v", err)
	}

	// Asynchronously write to Stdin.
	ch := make(chan iores)
	go asyncWrite(in, data, ch)

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

	// Close Stdin to signal to the program that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	// Wait for the program to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}
	return n, nil
}

// ReadFrom reads data from an io.Reader into the Portal.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	defer p.Reload()

	// Open a pipe to program stdin.
	in, err := p.StdinPipe()
	if err != nil {
		return 0, fmt.Errorf("portal: error during StdinPipe: %v", err)
	}

	// Asynchronously copy data from r into program stdin.
	ch := make(chan iores64)
	go asyncCopy(in, r, ch)

	// Start the program.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	// Receive results of asynchronous copy.
	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

	// Close program stdin to signal that we are done with it.
	if err = in.Close(); err != nil {
		return n, fmt.Errorf("portal: error while closing Stdin: %v", err)
	}

	// Wait for program to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}
	return n, nil
}
