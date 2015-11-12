package godash

import (
	"testing"
)

var (
	pln = Pln
	prf = Prf
)

func Test_pln(t *testing.T) {
	if l, e := pln("Hello World!!!"); l != 15 || e != nil {
		t.Error("pln() ... failed!")
	} else {
		t.Log("pln() ... ok!")
	}
}

func Test_prf(t *testing.T) {
	if l, e := prf("Hello %s!!!\n", "Golang"); l != 16 || e != nil {
		t.Error("prf() ... failed!")
	} else {
		t.Log("prf() ... ok!")
	}
}
