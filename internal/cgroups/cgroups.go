package cgroups

import (
	"os/exec"
	"testing"
)

func TestRunRequiresArguments(t *testing.T) {
	cmd := exec.Command("go", "run",
		"./cmd/byocontainer/", "invalid")
	out, _ := cmd.CombinedOutput()
	if len(out) == 0 {
		t.Error("expected error output for invalid command")
	}


}
