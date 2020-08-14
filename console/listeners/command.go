package listeners

import (
    event2 "github.com/mix-go/console/event"
    "github.com/mix-go/console/flag"
    "github.com/mix-go/console/process"
    "github.com/mix-go/event"
)

type CommandListener struct {
}

func (t *CommandListener) Events() []event.Event {
    return []event.Event{
        &event2.CommandBeforeExecuteEvent{},
    }
}

func (t *CommandListener) Process(e event.Event) {
    switch e.(type) {
    case *event2.CommandBeforeExecuteEvent:
        if flag.BoolMatch([]string{"d", "daemon"}, false) {
            process.Daemon()
        }
        break
    }
}
