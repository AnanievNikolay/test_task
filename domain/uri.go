package domain

import (
	"fmt"
	"strings"
)

//NewURI ...
func NewURI(_host, _fsym, _tsym string) *URI {
	return &URI{
		host: _host,
		fsym: _fsym,
		tsym: _tsym,
	}
}

//URI ...
type URI struct {
	host string
	fsym string
	tsym string
}

func (u *URI) String() string {
	return fmt.Sprintf("%v?fsym=%v&tsym=%v", u.host, strings.ToUpper(u.fsym), strings.ToLower(u.tsym))
}
