// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authentication

import (
	"github.com/xfali/neve-security/core/authority"
)

type Authentication interface {
	GetName() string

	GetAuthorities() []authority.GrantedAuthority

	GetCredentials() interface{}

	GetDetails() interface{}

	GetPrincipal() interface{}

	IsAuthenticated() bool

	SetAuthenticated(flag bool) error
}
