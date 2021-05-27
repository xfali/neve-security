// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package userdetails

import "github.com/xfali/neve-security/core/authority"

type User struct {
	Username             string
	Password             string
	IsAccountExpired     bool
	IsAccountLocked      bool
	IsDisabled           bool
	IsCredentialsExpired bool
	GrantedAuthority     []authority.GrantedAuthority
}

func (u *User) GetAuthorities() []authority.GrantedAuthority {
	return u.GrantedAuthority
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) IsAccountNonExpired() bool {
	return !u.IsAccountExpired
}

func (u *User) IsAccountNonLocked() bool {
	return !u.IsAccountLocked
}

func (u *User) IsCredentialsNonExpired() bool {
	return !u.IsCredentialsExpired
}

func (u *User) IsEnabled() bool {
	return !u.IsDisabled
}
