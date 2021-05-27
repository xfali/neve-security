// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package core

type Setter interface {
	Set(key string, value interface{})
}

// Bean注册配置，已支持的配置有：
// * bean.SetOrder(int) 配置bean注入顺序
type Opt func(setter Setter)
