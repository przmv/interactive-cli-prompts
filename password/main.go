package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

// PasswordPrompt asks for a string value using the label.
// The entered value will not be displayed on the screen
// while typing.
func PasswordPrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}

func main() {
	password := PasswordPrompt("What is your password?")
	fmt.Printf("Oh, I see! Your password is %q\n", password)
}
