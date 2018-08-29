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

	// Asynchronously read from Stdout to dst.
	ch := make(chan iores)
	go asyncRead(out, dst, ch)

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	// Receive results of read operation.
	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

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

	// Asynchronously copy data from Stdout into w.
	ch := make(chan iores64)
	go asyncCopy(w, out, ch)

	// Start Cmd.
	if err = p.Start(); err != nil {
		return 0, fmt.Errorf("portal: error while starting Cmd: %v", err)
	}

	// Receive results of copy operation.
	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

	// Wait for Cmd to exit.
	if err = p.Wait(); err != nil {
		return 0, fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}
	return n, nil
}
