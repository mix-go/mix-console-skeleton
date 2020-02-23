package main

import (
    "mix-skeleton/beans"
    "mix/src/bean"
    "mix/src/console"
)

// 应用清单
func Manifest() console.ApplicationDefinition {
    return console.ApplicationDefinition{
        AppName:    "app",
        AppVersion: "1.0.0",
        AppDebug:   true,
        Beans:      beans.Beans(),
        Commands: []console.CommandDefinition{
            {
                Name:        "foo",
                Reflect:     bean.NewReflect(FooCommand{}),
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
                Reflect:     bean.NewReflect(BarCommand{}),
                Description: "",
            },
        },
    }
}
