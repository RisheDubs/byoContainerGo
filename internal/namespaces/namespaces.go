package namespaces

import "syscall"

func DefaultFlags() uintptr {
	return syscall.CLONE_NEWUTS |
		syscall.CLONE_NEWPID |
		syscall.CLONE_NEWNS
}
