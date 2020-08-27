package commands

import (
    "github.com/mix-go/console"
    "github.com/mix-go/mix-console-skeleton/commands"
)

func init() {
    Commands = append(Commands,
        console.CommandDefinition{
            Name:    "wp",
            Usage:   "\tWorker pool demo",
            Command: &commands.WorkerPoolCommand{},
        },
        console.CommandDefinition{
            Name:  "wpd",
            Usage: "\tWorker pool daemon demo",
            Options: []console.OptionDefinition{
                {
                    Names: []string{"d", "daemon"},
                    Usage: "Run in the background",
                },
            },
            Command: &commands.WorkerPoolDaemonCommand{},
        },
    )
}
