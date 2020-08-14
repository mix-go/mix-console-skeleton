package commands

import (
    "fmt"
    "github.com/mix-go/console/flag"
)

type HelloCommand struct {
}

func (t *HelloCommand) Main() {
    name := flag.StringMatch([]string{"n", "name"}, "Xiao Ming")
    say := flag.StringMatch([]string{"say"}, "Hello, World!")
    fmt.Printf("%s: %s\n", name, say)
}
