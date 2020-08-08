package commands

import (
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
    "mix-skeleton/app/console/commands"
)

var (
    Commands []console.CommandDefinition
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "foo",
            Usage: "foo desc",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "your name",
                },
            },
            Reflect: bean.NewReflect(commands.FooCommand{}),
        }, console.CommandDefinition{
            Name:    "bar",
            Usage:   "bar desc",
            Reflect: bean.NewReflect(commands.BarCommand{}),
        },
    )
}
