package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/rishekesh/byoContainerGo/internal/namespaces"
)

func Parent(args []string) {
	cmd := exec.Command("/proc/self/exe",
		append([]string{"child"}, args...)...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: namespaces.DefaultFlags(),
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
