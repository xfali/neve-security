// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package userdetails

import (
	"github.com/xfali/neve-security/core/secerr"
)

type defaultChecker struct{}

func NewDefaultChecker() *defaultChecker {
	return &defaultChecker{}
}

func (c *defaultChecker) Check(details UserDetails) error {
	if !details.IsEnabled() {
		return secerr.UserAccountDisableError
	} else if !details.IsAccountNonExpired() {
		return secerr.UserAccountExpired
	} else if !details.IsAccountNonLocked() {
		return secerr.UserAccountLocked
	}
	return nil
}
