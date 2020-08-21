package beans

import "github.com/mix-go/bean"

var (
    Beans []bean.BeanDefinition
)

func Init() {
    dbInit()
    errorInit()
    eventInit()
    logInit()
    redisInit()
}
