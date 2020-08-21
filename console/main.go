package main

import (
    "console/globals"
    "console/manifest"
    "fmt"
    "github.com/jinzhu/configor"
    console2 "github.com/mix-go/console"
    "github.com/mix-go/console/argv"
    "github.com/mix-go/dotenv"
)

func init() {
    // Env
    if err := dotenv.Load(fmt.Sprintf("%s/../.env", argv.Program.Dir)); err != nil {
        panic(err)
    }
    // Conf
    if err := configor.Load(&globals.Config, fmt.Sprintf("%s/../conf/config.json", argv.Program.Dir)); err != nil {
        panic(err)
    }
    // manifest
    manifest.Init()
}

func main() {
    // App
    console2.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
}
