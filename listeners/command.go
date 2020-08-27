package listeners

import (
    "github.com/mix-go/mix-console-skeleton/globals"
    "github.com/mix-go/console"
    "github.com/mix-go/console/flag"
    "github.com/mix-go/console/process"
    "github.com/mix-go/event"
)

type CommandListener struct {
}

func (t *CommandListener) Events() []event.Event {
    return []event.Event{
        &console.CommandBeforeExecuteEvent{},
    }
}

func (t *CommandListener) Process(e event.Event) {
    switch e.(type) {
    case *console.CommandBeforeExecuteEvent:
        // 初始化全局对象
        globals.Init()

        // 设置守护
        if flag.Match("d", "daemon").Bool() {
            process.Daemon()
        }
        break
    }
}
