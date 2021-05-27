// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package authentication

type Provider interface {
	Authenticate(auth Authentication) (Authentication, error)
	Support(o interface{}) bool
}
