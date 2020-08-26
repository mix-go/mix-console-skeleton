package commands

import (
    "fmt"
    "github.com/mix-go/mix-skeleton/console/globals"
    "github.com/mix-go/mix-skeleton/console/models"
)

type DBCommand struct {
}

func (t *DBCommand) Main() {
    orm := globals.DB()
    var user models.User
    orm.Where("id = ?", 1).First(&user)
    fmt.Println(fmt.Sprintf("%+v", user))
}
