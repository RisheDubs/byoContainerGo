package container

import (
	"fmt"
	"os/exec"
	"syscall"
)

func Child(args []string) {
	setupFilesystem()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	
		
	}
}


