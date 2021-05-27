// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authority

type SimpleAuthority struct {
	role string
}

func NewSimpleAuthority(role string) *SimpleAuthority {
	return &SimpleAuthority{
		role: role,
	}
}

func (auth SimpleAuthority) GetAuthority() string {
	return auth.role
}
