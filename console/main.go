package main

import (
    "console/globals"
    "console/manifest"
    console2 "github.com/mix-go/console"
)

func main() {
    console2.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
    globals.Init()
}
