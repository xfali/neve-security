// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authority

type GrantedAuthority interface {
	GetAuthority() string
}
