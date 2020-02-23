package beans

import (
    "mix/src/bean"
    "net/http"
    "time"
)

func HttpBeans() []bean.BeanDefinition {
    return []bean.BeanDefinition{
        {
            Name:    "httpclient",
            Scope:   bean.SINGLETON,
            Reflect: bean.NewReflect(http.Client{}),
            Fields: bean.Fields{
                "Timeout": time.Duration(time.Second * 3),
            },
        },
    }
}
