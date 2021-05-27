// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package secerr

var (
	UserAccountDisableError = NewAuthenticationError("User Account is disabled. ")
	UserAccountExpired      = NewAuthenticationError("User Account is expired. ")
	UserAccountLocked       = NewAuthenticationError("User Account is locked. ")
)

type AuthenticationError struct {
	msg string `json:""`
}

func NewAuthenticationError(msg string) *AuthenticationError {
	return &AuthenticationError{
		msg: msg,
	}
}

func (e *AuthenticationError) Error() string {
	return e.msg
}
