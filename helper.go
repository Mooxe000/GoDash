package godash

import (
	"github.com/kr/pretty"
)

func Dd(obj interface{}) (int, error) {
	r, e := Prf("%# v\n", pretty.Formatter(obj))
	return r, e
}
