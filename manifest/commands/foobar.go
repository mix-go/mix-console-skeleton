package commands

import (
    "mix-skeleton/app/console/commands"
    "mix/src/bean"
    "mix/src/console"
)

// foo bar
func Foobar() []console.CommandDefinition {
    return []console.CommandDefinition{
        {
            Name:  "foo",
            Usage: "foo desc",
            Options: []console.CommandOption{
                {
                    Names: []string{"n", "name"},
                    Usage: "your name",
                },
            },
            Reflect: bean.NewReflect(commands.FooCommand{}),
        },
        {
            Name:    "bar",
            Usage:   "",
            Reflect: bean.NewReflect(commands.BarCommand{}),
        },
    }
}
