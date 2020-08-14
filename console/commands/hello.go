package commands

import (
    "fmt"
    "github.com/mix-go/console/flag"
)

type HelloCommand struct {
}

func (t *HelloCommand) Main() {
    name := flag.StringMatch("Xiao Ming", "n", "name")
    say := flag.StringMatch("Hello, World!", "say")
    fmt.Printf("%s: %s\n", name, say)
}
