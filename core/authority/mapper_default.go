// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authority

type defaultMapper struct {
}

func NewDefaultMapper() *defaultMapper {
	return &defaultMapper{}
}

func (m *defaultMapper) MapAuthorities(auth []GrantedAuthority) []GrantedAuthority {
	return auth
}
