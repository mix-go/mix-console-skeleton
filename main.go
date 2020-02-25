package main

import (
    "mix-skeleton/manifest"
    "mix/src/console"
)

var (
    App = console.NewApplication(manifest.Manifest())
)

// 主函数
func main() {
    App.Run()
}
