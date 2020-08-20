package main

import (
    "console/manifest"
    console2 "github.com/mix-go/console"
    "github.com/mix-go/dotenv"
)

func main() {
    // Env
    _ = dotenv.Load("../.env")
    // App
    console2.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
}
