// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authentication

import (
	"github.com/xfali/neve-security/core"
	"github.com/xfali/neve-security/core/secerr"
	"github.com/xfali/neve-utils/reflection"
	"github.com/xfali/xlog"
)

const (
	KeyManagerAddProvider = "authentication.manager.provider.add"
)

type defaultManager struct {
	Providers []Provider `inject:""`

	logger xlog.Logger
}

func NewManager(opts ...core.Opt) *defaultManager {
	ret := &defaultManager{
		logger: xlog.GetLogger(),
	}

	for _, opt := range opts {
		opt(ret)
	}

	return ret
}

func (m *defaultManager) Authenticate(auth Authentication) (Authentication, error) {
	var lastErr *secerr.AuthenticationError
	for _, provider := range m.Providers {
		if provider.Support(auth) {
			if m.logger.DebugEnabled() {
				m.logger.Debugf("auth: %s use provider: %s\n", auth.GetName(), reflection.GetObjectName(provider))
			}

			dest, err := provider.Authenticate(auth)
			if err != nil {
				if e, ok := err.(*secerr.AuthenticationError); ok {
					lastErr = e
				} else {
					return nil, err
				}
			} else {
				return dest, nil
			}
		}
	}
	return nil, lastErr
}

func (m *defaultManager) Set(key string, value interface{}) {
	switch key {
	case KeyManagerAddProvider:
		m.Providers = append(m.Providers, value.([]Provider)...)
	}
}

func ManagerOptAddProviders(providers ...Provider) core.Opt {
	return func(setter core.Setter) {
		setter.Set(KeyManagerAddProvider, providers[:])
	}
}
