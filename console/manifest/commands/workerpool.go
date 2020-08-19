package commands

import (
    "console/commands"
    "github.com/mix-go/bean"
    "github.com/mix-go/console"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:  "wp",
            Usage: "\tWorker pool demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"d", "daemon"},
                    Usage: "Run in the background",
                },
            },
            Reflect: bean.NewReflect(commands.HelloCommand{}),
        },
    )
}
