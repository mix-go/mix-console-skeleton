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
            Name:        "foo",
            Reflect:     bean.NewReflect(commands.FooCommand{}),
            Description: "foo desc",
            Options: []console.CommandOption{
                {
                    Names:       []string{"n", "name"},
                    Description: "your name",
                },
            },
        },
        {
            Name:        "bar",
            Reflect:     bean.NewReflect(commands.BarCommand{}),
            Description: "",
        },
    }
}
