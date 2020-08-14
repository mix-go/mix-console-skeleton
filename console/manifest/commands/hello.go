package commands

import (
    "console/commands"
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
)

var (
    Commands []console.CommandDefinition
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "hello",
            Usage: "\tEcho demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"n", "name"},
                    Usage: "Your name",
                },
                {
                    Names: []string{"say"},
                    Usage: "\tSay ...",
                },
                {
                    Names: []string{"d", "daemon"},
                    Usage: "Run in the background",
                },
            },
            Reflect: bean.NewReflect(commands.HelloCommand{}),
        },
    )
}
