package godash

import (
	"testing"
)

func Test_ToInterfaces(t *testing.T) {
	FuncName := "GoInterface.ToInterfaces()"
	ss := []string{"1", "2", "3", "4", "5"}
	gi := GdInterface{ss}
	Pgi := &gi
	ins := Pgi.ToInterfaces()

	// dd(typeof(ins) == "[]interface {}")
	// dd(IsSlice(ins))

	if IsSlice(ins) && typeof(ins) == "[]interface {}" {
		for i, v := range ins {
			if ss[i] != v {
				t.Error(FuncName + " ... failed!")
			}
		}
	} else {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + "... ok!")
}
