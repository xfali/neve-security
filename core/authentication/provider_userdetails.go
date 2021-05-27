// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authentication

import (
	"github.com/xfali/neve-security/core"
	"github.com/xfali/neve-security/core/authority"
	"github.com/xfali/neve-security/core/userdetails"
)

const (
	KeyProviderPreChecker      = "authentication.provider.prechecker.set"
	KeyProviderPostChecker     = "authentication.provider.postchecker.set"
	KeyProviderAuthorityMapper = "authentication.provider.authority.mapper.set"
)

type UserDetailsAuthenticationChecker interface {
	Retrieve(username string, authentication *UsernamePasswordAuthentication) (userdetails.UserDetails, error)
	AdditionalAuthenticationCheck(username string, authentication *UsernamePasswordAuthentication) error
}

type UserDetailsAuthenticationProvider struct {
	UserDetailsAuthenticationChecker

	preChecker  userdetails.Checker
	postChecker userdetails.Checker
	mapper      authority.Mapper
}

func NewUserDetailsAuthenticationProvider(opts ...core.Opt) *UserDetailsAuthenticationProvider {
	ret := &UserDetailsAuthenticationProvider{
		preChecker:  userdetails.NewDefaultChecker(),
		postChecker: userdetails.NewDefaultChecker(),
		mapper:      authority.NewDefaultMapper(),
	}

	for _, opt := range opts {
		opt(ret)
	}

	return ret
}

func (p *UserDetailsAuthenticationProvider) Set(key string, value interface{}) {
	switch key {
	case KeyProviderPreChecker:
		p.preChecker = value.(userdetails.Checker)
	case KeyProviderPostChecker:
		p.postChecker = value.(userdetails.Checker)
	case KeyProviderAuthorityMapper:
		p.mapper = value.(authority.Mapper)
	}
}

func (p *UserDetailsAuthenticationProvider) Authenticate(auth Authentication) (Authentication, error) {
	if userPwAuth, ok := auth.(*UsernamePasswordAuthentication); ok {
		username := userPwAuth.GetName()
		userdetails, err := p.Retrieve(username, userPwAuth)
		if err != nil {
			return nil, err
		}
		err = p.preChecker.Check(userdetails)
		if err != nil {
			return nil, err
		}
		err = p.AdditionalAuthenticationCheck(username, userPwAuth)
		if err != nil {
			return nil, err
		}

		err = p.postChecker.Check(userdetails)
		if err != nil {
			return nil, err
		}

		return p.createSuccessAuthentication(auth, userdetails, userdetails), nil
	} else {
		panic("Support UsernamePasswordAuthentication only")
	}
}

func (p *UserDetailsAuthenticationProvider) createSuccessAuthentication(auth Authentication, principal interface{}, user userdetails.UserDetails) Authentication {
	return NewUsernamePasswordAuthentication(auth.GetCredentials(), principal, p.mapper.MapAuthorities(user.GetAuthorities()))
}

func (p *UserDetailsAuthenticationProvider) Support(o interface{}) bool {
	panic("Not support")
}

func ProviderOptSetPreChecker(checker userdetails.Checker) core.Opt {
	return func(setter core.Setter) {
		setter.Set(KeyProviderPreChecker, checker)
	}
}

func ProviderOptSetPostChecker(checker userdetails.Checker) core.Opt {
	return func(setter core.Setter) {
		setter.Set(KeyProviderPostChecker, checker)
	}
}

func ProviderOptSetAuthorityMapper(mapper authority.Mapper) core.Opt {
	return func(setter core.Setter) {
		setter.Set(KeyProviderAuthorityMapper, mapper)
	}
}
