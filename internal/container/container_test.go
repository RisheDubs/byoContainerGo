package container_test

import (
	"os/exec"
	"testing"
)

func TestBinaryExistsAndBuilds(t *testing.T) {
	cmd := exec.Command("go", "build", "./...")
	if err := cmd.Run(); err != nil {
		t.Errorf("project failed to build: %v", err)
	}
}

func TestRunRequiresArguments(t *testing.T) {
	cmd := exec.Command("go", "run",
		"./cmd/byocontainer/", "invalid")
	out, _ := cmd.CombinedOutput()
	if len(out) == 0 {
		t.Error("expected error output for invalid command")
	}
}
