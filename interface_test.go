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

func Test_ToStrings(t *testing.T) {
	FuncName := "GoInterface.ToStrings()"

  ss := []string{"1", "2", "3", "4", "5"}
	gi := GdInterface{ss}
	Pgi := &gi
	ins := Pgi.ToInterfaces()

  gis := GdInterfaces{ins}
  Pgis := &gis
  nss := Pgis.ToStrings()

  for i, v := range nss {
    if v != ss[i] {
      t.Error(FuncName + " ... failed!")
    }
  }

	t.Log(FuncName + "... ok!")
}
