package globals

import (
    "github.com/mix-go/console"
    "github.com/mix-go/logrus"
)

func GetLogger() *logrus.Logger {
    return console.Context().Get("logger").(*logrus.Logger)
}
