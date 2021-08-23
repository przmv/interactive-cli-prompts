## Tidal Migrations 💓 CLI applications

Do you like CLI applications? We love them! At Tidal Migrations we use full-featured GUI IDEs and editors like VS Code and Emacs but also `vim` and `git` running in our terminals. Every day we use `bash`, `awk`, `sed` and lots of other CLI tools and apps for work and fun. Also, we like to develop [CLI apps](https://get.tidal.sh/) and with this post, we're going to show you how to implement different interactive prompts for your CLI apps written in Go.

## Passing data to CLI apps

Oftentimes CLI applications don't just work by themselves, but some process or operation is required on the information or data.

There are different ways to pass data to command line applications. Using flags, environment variables, file names as CLI arguments or reading from standard input is quite common and is pretty easy to implement using just the [standard Go library](https://pkg.go.dev/std). Using interactive prompts can _spice up_ your CLI application and improve the overall UX.

Let's get started!

## How to implement text input prompt

The basic text input prompt is easy to implement. Just read from standard input until the [new line character](https://en.wikipedia.org/wiki/Newline) (`\n`):

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	name := StringPrompt("What is your name?")
	fmt.Printf("Hello, %s!\n", name)
}
```

{% asciinema 430306 %}

## How to implement password input prompt

Password prompts are similar to text input prompts, except the user's typed input should be hidden:

```go
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
```

{% asciinema 430310 %}

## How to implement Yes/No prompt

For Yes/No prompts we're going to create an infinite loop to keep asking until the user answers yes or no:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// YesNoPrompt asks yes/no questions using the label.
func YesNoPrompt(label string, def bool) bool {
	choices := "Y/n"
	if !def {
		choices = "y/N"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "%s (%s) ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return def
		}
		s = strings.ToLower(s)
		if s == "y" || s == "yes" {
			return true
		}
		if s == "n" || s == "no" {
			return false
		}
	}
}

func main() {
	ok := YesNoPrompt("Dev.to is awesome!", true)
	if ok {
		fmt.Println("Agree!")
	} else {
		fmt.Println("Huh?")
	}
}
```

{% asciinema 430311 %}

## How to implement interactive checkboxes

To create an interactive multi-select prompt we're going to use an awesome [`survey`](https://github.com/AlecAivazis/survey) package:

```go
package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Checkboxes(label string, opts []string) []string {
	res := []string{}
	prompt := &survey.MultiSelect{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func main() {
	answers := Checkboxes(
		"Which are your favourite programming languages?",
		[]string{
			"C",
			"Python",
			"Java",
			"C++",
			"C#",
			"Visual Basic",
			"JavaScript",
			"PHP",
			"Assembly Language",
			"SQL",
			"Groovy",
			"Classic Visual Basic",
			"Fortran",
			"R",
			"Ruby",
			"Swift",
			"MATLAB",
			"Go",
			"Prolog",
			"Perl",
		},
	)
	s := strings.Join(answers, ", ")
	fmt.Println("Oh, I see! You like", s)
}
```

{% asciinema 430317 %}

## Caveats and workarounds

If you pipe some input data to your interactive CLI app, the prompts will read that data:

```terminal
$ echo "Petr" | go run main.go
What is your name? Hello, Petr!
```

Sometimes such behavior is acceptable, but sometimes not. To check if the terminal is interactive let's use [`term.IsTerminal`](https://pkg.go.dev/golang.org/x/term#IsTerminal) function:

```go
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
```

```terminal
$ echo "Hello" | go run main.go
Terminal is not interactive! Consider using flags or environment variables!

$ go run main.go
Terminal is interactive! You're good to use prompts!
```

## Libraries

As you can see, it's pretty easy to implement basic interactive prompts, but for complex ones it's better to use some Go packages from the community:

{% github AlecAivazis/survey %}
{% github Songmu/prompter %}
{% github manifoldco/promptui %}

## Conclusion

That's it! We hope you liked it! Code examples are available on [GitHub](https://github.com/tidalmigrations/interactive-cli-prompts).

If you're interested in CLI applications development in Go and we — [Tidal Migrations](https://tidalmigrations.com/) — are hiring! Please check our [Careers](https://tidalmigrations.com/careers/) page!

Long live the command line!
