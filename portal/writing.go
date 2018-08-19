package portal

import (
	"fmt"
	"io"
)

// Write allows for the writing of data into a command's standard input.
func (p *Portal) Write(src []byte) (n int, err error) {
	p.Restore()
	in, err := p.StdinPipe()
	if err != nil {
		return 0, stdinPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = in.Write(src); err != nil {
		return n, fmt.Errorf("portal: could not write to Stdin: %v", err)
	}
	if err = in.Close(); err != nil {
		return n, closeStdinErr(err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, err
}

// ReadFrom allows for the piping of data from a io.Writer into a command's
// standard output.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	p.Restore()
	in, err := p.StdinPipe()
	if err != nil {
		return 0, stdinPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = io.Copy(in, r); err != nil {
		return n, fmt.Errorf("portal: failed to write to Stdin: %v", err)
	}
	if err = in.Close(); err != nil {
		return n, closeStdinErr(err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, err
}
