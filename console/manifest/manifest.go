package manifest

import (
    "github.com/mix-go/console"
    "mix-skeleton/manifest/beans"
    "mix-skeleton/manifest/commands"
)

var (
    ApplicationDefinition console.ApplicationDefinition
)

func init() {
    ApplicationDefinition = console.ApplicationDefinition{
        AppName:    "app",
        AppVersion: "1.0.0",
        AppDebug:   true,
        Beans:      beans.Beans,
        Commands:   commands.Commands,
    }
}
