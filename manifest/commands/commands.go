package commands

import (
    "mix/src/console"
)

// 全部命令配置
func Commands() []console.CommandDefinition {
    cmds := []console.CommandDefinition{}
    cmds = append(cmds, Foobar()...)
    return cmds
}
