package main

import(
    "fmt"
    "os"

    "github.com/rishekesh/byoContainerGo/internal/container"
)

func main() {
	if len(os.Args) < 2{
		fmt.Fprintln(os.Stderr, "usage: byocontainer run <cmd>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		container.Parent(os.Args[2:])
	case "child":
		container.Child(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
