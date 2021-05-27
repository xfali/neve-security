// Copyright (C) 2019-2021, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package userdetails

type Checker interface {
	Check(details UserDetails) error
}
