package console

import (
    "github.com/mix-go/console"
    "mix-skeleton/manifest"
)

func main() {
    console.NewApplication(manifest.ApplicationDefinition, "eventDispatcher", "error").Run()
}
