package portal

import "io"

// Read reads len(dst) of data from the Portal into dst.
func (p *Portal) Read(dst []byte) (n int, err error) {
	defer p.Reload()

	// Open an pipe to Stdout.
	out, err := p.stdoutPipe()
	if err != nil {
		return 0, err
	}

	// Asynchronously read from Stdout to dst.
	ch := make(chan iores)
	go asyncRead(out, dst, ch)

	// Start Cmd.
	if err = p.start(); err != nil {
		return 0, err
	}

	// Receive results of read operation.
	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

	// Wait for Cmd to complete.
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

	// Asynchronously copy data from Stdout into w.
	ch := make(chan iores64)
	go asyncCopy(w, out, ch)

	// Start Cmd.
	if err = p.start(); err != nil {
		return 0, err
	}

	// Receive results of copy operation.
	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	n = res.n

	// Wait for Cmd to exit.
	return n, p.wait()
}
