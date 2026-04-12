package namespaces_test

import (
	"testing"
	"syscall"

	"github.com/rishekesh/byoContainerGo/internal/namespaces"
)

func TestDefaultFlagsIncludesUTS(t *testing.T) {
	flags := namespaces.DefaultFlags()
	if flags&syscall.CLONE_NEWUTS == 0 {
		t.Error("expected CLONE_NEWUTS flag to be set")
	}
}

func TestDefaultFlagsIncludesPID(t *testing.T) {
	flags := namespaces.DefaultFlags()
	if flags&syscall.CLONE_NEWPID == 0 {
		t.Error("expected CLONE_NEWPID flag to be set")
	}
}

func TestDefaultFlagsIncludesMNT(t *testing.T) {
	flags := namespaces.DefaultFlags()
	if flags&syscall.CLONE_NEWNS == 0 {
		t.Error("expected CLONE_NEWNS flag to be set")
	}
}
