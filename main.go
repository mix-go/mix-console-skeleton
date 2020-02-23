package main

import (
    "mix/src/console"
)

type FooCommand struct {
}
type BarCommand struct {
}

func main() {
    app := console.NewApplication(Manifest())
    app.Run()
}
