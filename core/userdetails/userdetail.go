// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package userdetails

import (
	"github.com/xfali/neve-security/core/authority"
)

type UserDetails interface {
	GetAuthorities() []authority.GrantedAuthority

	GetPassword() string
	GetUsername() string

	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool

	IsEnabled() bool
}
