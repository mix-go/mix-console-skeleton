package manifest

import (
    "mix-skeleton/manifest/beans"
    "mix-skeleton/manifest/commands"
    "mix/src/console"
)

// 应用清单配置
func Manifest() console.ApplicationDefinition {
    return console.ApplicationDefinition{
        AppName:    "app",
        AppVersion: "1.0.0",
        AppDebug:   true,
        Beans:      beans.Beans(),
        Commands:   commands.Commands(),
    }
}
