// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authority

type Mapper interface {
	MapAuthorities(auth []GrantedAuthority) []GrantedAuthority
}
