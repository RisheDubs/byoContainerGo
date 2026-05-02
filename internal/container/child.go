package container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Child(args []string) {
	setupFilesystem()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		
	}
}


func must(err error) {
	if err != nil {
		panic(err)
	}
}
