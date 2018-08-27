package portal_test

import (
	"github.com/steven-xie/glip/portal"
	"os/exec"
	"testing"
)

// Global Portal instances to be used by the testing package.
var (
	catp  = portal.New("cat")
	echop = portal.New("echo", TestPhrase)
)

func TestPortal(t *testing.T) {
	var (
		echoCmd    = exec.Command("echo", TestPhrase)
		echoPortal = portal.NewFrom(echoCmd)
	)

	if !echoPortal.IsReady() {
		t.Fatal("Expected prepared portal to be ready for execution.")
	}

	if err := echoPortal.Run(); err != nil {
		t.Fatalf("Got an error while running portal: %v", err)
	}

	if echoPortal.IsReady() {
		t.Fatal("Expected portal to require reloading before it is ready after " +
			"running.")
	}
}
