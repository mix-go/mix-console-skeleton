package beans

import (
    "github.com/mix-go/bean"
    "net/http"
    "time"
)

var (
    Beans []bean.BeanDefinition
)

func init() {
    Beans = append(Beans,
        bean.BeanDefinition{
            Name:    "httpclient",
            Scope:   bean.SINGLETON,
            Reflect: bean.NewReflect(http.Client{}),
            Fields: bean.Fields{
                "Timeout": time.Duration(time.Second * 3),
            },
        },
    )
}
