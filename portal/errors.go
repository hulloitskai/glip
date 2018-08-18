package portal

import "fmt"

func stdoutPipeErr(err error) error {
	return fmt.Errorf("portal: could not connect to Stdout: %v", err)
}

func stdinPipeErr(err error) error {
	return fmt.Errorf("portal: could not connect to Stdin: %v", err)
}

func closeStdinErr(err error) error {
	return fmt.Errorf("portal: could not close Stdin: %v", err)
}

func startErr(err error) error {
	return fmt.Errorf("portal: could not start command: %v", err)
}

func waitErr(err error) error {
	return fmt.Errorf("portal: could not wait for command completion: %v", err)
}
