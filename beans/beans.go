package beans

import "mix/src/bean"

func Beans() []bean.BeanDefinition {
    bs := []bean.BeanDefinition{}
    bs = append(bs, HttpBeans()...)
    return bs
}
