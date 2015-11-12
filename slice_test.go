package godash

import (
	"testing"
)

func Test_String_Pop(t *testing.T) {
	typeof := TypeOf

	ms := MxSlice{"1", "2", "3", "4", "5"}

	if ms.Len() != 5 {
		t.Error("MxSlice.Len() ... failed!")
	}

	if ms.At(3) != "4" {
		t.Error("MxSlice.At() ... failed!")
	}

	v, _ := ms.Slice(1, 3)
	if typeof(v) != "util.MxSlice" && v.Len() != 2 {
		t.Error("MxSlice.Slice() ... failed!")
	}

	a := []string{"2", "3"}
	b := v.OutputStrings()
	for ii, i := range a {
		for ij, j := range b {
			if ii == ij && i != j {
				t.Error("MxSlice.Slice() ... failed!")
			}
		}
	}

	wms := WMxSliceItf{ms}
	c, l := wms.Pop()
	if c != "5" {
		t.Error("Pop() ... failed!")
	}
	d := []string{"1", "2", "3", "4"}
	e := l.OutputStrings()
	if typeof(l) != "util.MxSlice" && l.Len() != 4 {
		t.Error("Pop() ... failed!")
	}
	for ig, g := range d {
		for ih, h := range e {
			if ig == ih && g != h {
				t.Error("MxSlice.Slice() ... failed!")
			}
		}
	}

	t.Log("Pop_String() ... ok!")
}

// func Test_String_Push(t *testing.T) {
//
// 	t.Log("Push_String() ... ok!")
// }
