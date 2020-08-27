package commands

import (
    "github.com/mix-go/console"
    "github.com/mix-go/mix-console-skeleton/commands"
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
            },
            Command: &commands.HelloCommand{},
        },
    )
}
