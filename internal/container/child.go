package container

import (
	"fmt"
	
	
)

func Child(args []string) {
	setupFilesystem()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	

	
		
	}
}


