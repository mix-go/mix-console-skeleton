package main

import (
    "github.com/mix-go/console"
    "mix-skeleton/manifest"
)

func main() {
    console.NewApplication(manifest.ApplicationDefinition).Run()
}
