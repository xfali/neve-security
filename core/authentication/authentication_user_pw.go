// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authentication

import (
	"fmt"
	"github.com/xfali/neve-security/core/authority"
	"github.com/xfali/neve-security/core/userdetails"
	"sync/atomic"
)

type BaseAuthentication struct {
	Authentication
	authorities   []authority.GrantedAuthority
	authenticated int32
	details       interface{}
}

type UsernamePasswordAuthentication struct {
	BaseAuthentication

	credentials interface{}
	principal   interface{}
}

func NewUsernamePasswordAuthentication(credentials, principal interface{}, authorities []authority.GrantedAuthority) *UsernamePasswordAuthentication {
	ret := &UsernamePasswordAuthentication{
		credentials: credentials,
		principal:   principal,
	}
	if len(authorities) == 0 {
		ret.SetAuthenticated(false)
	} else {
		ret.SetAuthenticated(true)
		ret.authorities = authorities
	}
	return ret
}

func (auth *BaseAuthentication) GetName() string {
	p := auth.GetPrincipal()
	if d, ok := p.(userdetails.UserDetails); ok {
		return d.GetUsername()
	} else {
		return fmt.Sprintf("%v", p)
	}
}

func (auth *BaseAuthentication) GetAuthorities() []authority.GrantedAuthority {
	return auth.authorities
}

func (auth *BaseAuthentication) IsAuthenticated() bool {
	return atomic.LoadInt32(&auth.authenticated) == 1
}

func (auth *BaseAuthentication) SetAuthenticated(flag bool) error {
	if flag {
		atomic.StoreInt32(&auth.authenticated, 1)
	} else {
		atomic.StoreInt32(&auth.authenticated, 0)
	}
	return nil
}

func (auth *BaseAuthentication) GetDetails() interface{} {
	return auth.details
}

func (auth *BaseAuthentication) SetDetails(details interface{}) {
	auth.details = details
}

func (auth *UsernamePasswordAuthentication) GetCredentials() interface{} {
	return auth.credentials
}

func (auth *UsernamePasswordAuthentication) GetPrincipal() interface{} {
	return auth.principal
}
