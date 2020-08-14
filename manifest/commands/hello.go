package commands

import (
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
    "mix-skeleton/console/commands"
)

var (
    Commands []console.CommandDefinition
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "hello",
            Usage: "Echo demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "your name",
                },
                {
                    Names: []string{"say"},
                    Usage: "Say ...",
                },
            },
            Reflect: bean.NewReflect(commands.HelloCommand{}),
        },
    )
}
