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
		echoPortal = portal.MakeFrom(echoCmd)
	)

	if echoPortal.Path != "" {
		t.Fatalf("Expected new portal to begin with an empty Cmd field (with an "+
			"empty Path), but found echoPortal.Cmd to be: %#v", echoPortal.Cmd)
	}

	echoPortal.Prepare()
	if !echoPortal.IsReady() {
		t.Fatal("Expected prepared portal to be ready for execution.")
	}
}
