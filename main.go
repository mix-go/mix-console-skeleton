package main

import (
    "mix-skeleton/manifest"
    "mix/src/console"
)

// 主函数
func main() {
    app := console.NewApplication(manifest.Manifest())
    app.Run()
}
