package error

import (
	"fmt"
)

type String interface {
	String() string
}

// type E interface {
// 	error | String
// }

type Error struct {
	Desc string
	E    any
}

func (e Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Desc, e.E)
}
