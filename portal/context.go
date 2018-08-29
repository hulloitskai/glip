package portal

import (
	"fmt"
	"io"
)

func (p *Portal) stdinPipe() (io.WriteCloser, error) {
	in, err := p.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("portal: error during StdinPipe: %v", err)
	}
	return in, nil
}

func (p *Portal) stdoutPipe() (io.ReadCloser, error) {
	in, err := p.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("portal: error during StdoutPipe: %v", err)
	}
	return in, nil
}

func (p *Portal) start() error {
	if err := p.Start(); err != nil {
		return fmt.Errorf("portal: error while starting Cmd: %v", err)
	}
	return nil
}

func (p *Portal) wait() error {
	if err := p.Wait(); err != nil {
		return fmt.Errorf("portal: error while waiting for Cmd to exit: %v", err)
	}
	return nil
}

func (p *Portal) run() error {
	if err := p.Run(); err != nil {
		return fmt.Errorf("portal: error while running Cmd: %v", err)
	}
	return nil
}
