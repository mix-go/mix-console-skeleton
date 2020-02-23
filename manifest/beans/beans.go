package beans

import "mix/src/bean"

// 全部依赖注入配置
func Beans() []bean.BeanDefinition {
    bs := []bean.BeanDefinition{}
    bs = append(bs, HttpBeans()...)
    return bs
}
