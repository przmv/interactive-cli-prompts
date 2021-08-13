package main

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func main() {
	if term.IsTerminal(int(syscall.Stdin)) {
		fmt.Println("Terminal is interactive! You're good to use prompts!")
	} else {
		fmt.Println("Terminal is not interactive! Consider using flags or environment variables!")
	}
}
