package godash

import (
	"testing"
)

func getGdListIns() (GdList, GdList) {
	var ssgl, iigl GdList

	ssmi := GdInterface{[]string{"1", "2", "3", "4", "5"}}
	ssgl.Value = ssmi.ToInterfaces()

	iimi := GdInterface{[]int{1, 2, 3, 4, 5}}
	iigl.Value = iimi.ToInterfaces()

	return ssgl, iigl
}

func Test_New(t *testing.T) {
	FuncName := "GdLists.New()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	// dd(Pssgl.Value)
	// dd(Pssgl.list)

	if Pssgl.list.Len() != 5 {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

// func Test_Sync(t *testing.T) {
//   FuncName := "GdLists.Sync()"
//   t.Log(FuncName + " ... ok!")
// }

func Test_Len(t *testing.T) {
	FuncName := "GdLists.Len()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	l := Pssgl.Len()
	a := len(Pssgl.Value)
	b := Pssgl.list.Len()

	// dd(v)
	// dd(a)
	// dd(b)

	if a != l || b != l {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_At(t *testing.T) {
	FuncName := "GdLists.At()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	l := Pssgl.Len()
	for i := 0; i < l; i++ {
		v := Pssgl.At(i).Value
		// dd(v)
		if v != Pssgl.Value[i] {
			t.Error(FuncName + " ... failed!")
		}
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Setter(t *testing.T) {
	FuncName := "GdLists.Setter()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	var v interface{} = "100"
	// dd(v)
	Pssgl.Setter(2, v)

	// dd(Pssgl.Value[2])

	if Pssgl.Value[2] != "100" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Remove(t *testing.T) {
	FuncName := "GdLists.Remove()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	Pssgl.Remove(2).Sync()

	// dd(Pssgl.Value)

	if Pssgl.Len() != 4 || Pssgl.Value[2] != "4" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Remove_Sync(t *testing.T) {
	FuncName := "GdLists.Remove() AND GdLists.Sync()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	v := Pssgl.Remove(2)

	l := Pssgl.Len()
	a := len(Pssgl.Value)
	b := Pssgl.list.Len()
	// dd(l)
	// dd(a)
	// dd(b)

	if l != b || a-b != 1 {
		t.Error(FuncName + " ... failed!")
	}

	v.Sync()

	ls := Pssgl.Len()
	as := len(Pssgl.Value)
	// bs := Pssgl.list.Len()
	// dd(ls)
	// dd(as)
	// dd(bs)

	if as != ls || ls != 4 || Pssgl.Getter(2) != "4" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_iCheck(t *testing.T) {
	FuncName := "GdLists.iCheck()"

	// i, l := 0, 5

	// -- i < 0
	// i = -5
	// -- i >= l
	// i = 10
	// -- 0 <= i < l
	// i = 2
	// i = 0
	// i = 4

	// dd(iCheck(i, l))

	t.Log(FuncName + " ... ok!")
}

func Test_seCheck(t *testing.T) {
	FuncName := "GdLists.seCheck()"

	// -- a == b == 0
	// s, e, l := 0, 0, 5       // 0, 0

	// -- a < 0 && b < 0
	// s, e = -5, -10           // 0, 0
	// -- a >=l && b >= l
	// s, e = 5, 10             // 4, 4
	// -- a < 0 && b >= l
	// s, e = -5, 10            // 0, 4
	// -- a < 0 && b >= 0
	// s, e = -5, 2             // 0, 2
	// -- a >= 0 && b >= l
	// s, e = 2, 10             // 2, 4

	// -- a == b == l - 1
	// s, e = 4, 4              // 4, 4
	// -- a == 0 && b == l - 1
	// s, e = 0, 4              // 0, 4
	// -- 0 < a < b < l
	// s, e = 1, 3              // 1, 3
	// -- a = 0
	// s, e = 0, 2              // 0, 2
	// -- b = l - 1
	// s, e = 2, 4              // 2, 4

	// a, b := seCheck(s, e, l)
	// dd(a)
	// dd(b)

	t.Log(FuncName + " ... ok!")
}

func Test_Slice(t *testing.T) {
	FuncName := "GdLists.Slice()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	// -- a == b == 0
	// s, e, _ := 0, 0, 5       // 1

	// -- a == b == l - 1
	// s, e = 4, 4              // 5
	// -- a == 0 && b == l - 1
	// s, e = 0, 4              // 1, 2, 3, 4, 5
	// -- 0 < a < b < l
	// s, e = 1, 3              // 2, 3, 4
	// a = 0
	// s, e = 0, 2              // 1, 2, 3
	// b = l - 1
	// s, e = 2, 4              // 3, 4, 5

	// v := Pssgl.Slice(s, e)

	// dd(v.Value)

	t.Log(FuncName + " ... ok!")
}

func Test_Chunk_1(t *testing.T) {
	FuncName := "GdLists.Chunk()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	////////////////////////////////
	// Chunk(0/5)
	////////////////////////////////
	gls_0 := Pssgl.Chunk(0)
	gls_5 := Pssgl.Chunk(5)

	// pln(len(gls_0))
	// pln(len(gls_5))

	if len(gls_0) != 1 || len(gls_5) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	// dd(gls_0[0].Value)
	// dd(gls_5[0].Value)

	if gls_0[0].Len() != len(gls_0[0].Value) || gls_5[0].Len() != 5 {
		t.Error(FuncName + " ... failed!")
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(1)
	////////////////////////////////
	gls_1 := Pssgl.Chunk(1)

	if len(gls_1) != 5 {
		t.Error(FuncName + " ... failed!")
	}

	for _, v := range gls_1 {
		// dd(v.Value)
		if len(v.Value) != 1 {
			t.Error(FuncName + " ... failed!")
		}
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(2)
	////////////////////////////////
	gls_2 := Pssgl.Chunk(2)

	is := []int{2, 2, 1}
	for i, v := range gls_2 {
		// dd(v.Value)
		if v.Len() != is[i] {
			t.Error(FuncName + " ... failed!")
		}
	}
	////////////////////////////////

	////////////////////////////////
	// Chunk(3)
	////////////////////////////////
	gls_3 := Pssgl.Chunk(3)

	ii := 0
	for _, v := range gls_3 {
		// dd(v.Value)
		ii = ii + v.Len()
	}

	if ii != 5 {
		t.Error(FuncName + " ... failed!")
	}

	////////////////////////////////
	// Chunk(4)
	////////////////////////////////
	gls_4 := Pssgl.Chunk(4)

	if len(gls_4) != 2 {
		t.Error(FuncName + " ... failed!")
	}

	// dd(gls_4[0].Value)
	// dd(gls_4[1].Value)

	if len(gls_4[0].Value) != 4 {
		t.Error(FuncName + " ... failed!")
	}
	if len(gls_4[1].Value) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_Chunk_2(t *testing.T) {
	FuncName := "GdLists.Chunk()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	////////////////////////////////
	// Chunk(0, 0, 0)
	////////////////////////////////
	gls_0 := Pssgl.Chunk(0, 0, 0)

	if len(gls_0) != 1 {
		t.Error(FuncName + " ... failed!")
	}

	for _, v := range gls_0 {
		// dd(v.Value)
		if len(v.Value) != 5 {
			t.Error(FuncName + " ... failed!")
		}
	}

	////////////////////////////////
	// Chunk(1, 1, 1)
	////////////////////////////////
	gls_1 := Pssgl.Chunk(1, 1, 1)

	if len(gls_1) != 4 {
		t.Error(FuncName + " ... failed!")
	}

	for i, v := range gls_1 {
		// dd(v.Value)
		if i <= 2 {
			if len(v.Value) != 1 {
				t.Error(FuncName + " ... failed!")
			}
		} else {
			if i != len(gls_1)-1 || len(v.Value) != 2 {
				t.Error(FuncName + " ... failed!")
			}
		}
	}

	////////////////////////////////
	// TODO Chunk
	////////////////////////////////
	// gls := Pssgl.Chunk(1, 3, 8)
	// gls := Pssgl.Chunk(3, 0, 1)
	// gls := Pssgl.Chunk(0, 2, 5)
	// gls := Pssgl.Chunk(2, 1, 2)
	// for _, v := range gls {
	//   dd(v.Value)
	// }

	t.Log(FuncName + " ... ok!")
}

func Test_Concat(t *testing.T) {
	FuncName := "GdLists.Concat()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	Pssgl.Concat(Pssgl.Chunk(2))

	// dd(Pssgl.Value)

	if Pssgl.Len() != 5 || Pssgl.Value[2] != "3" {
		t.Error(FuncName + " ... failed!")
	}

	t.Log(FuncName + " ... ok!")
}

func Test_RemoveList(t *testing.T) {
	FuncName := "GdLists.RemoveList()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	// -- a == b == 0
	// s, e, _ := 0, 0, 5        // 2, 3, 4, 5

	// -- a == b == l - 1
	// s, e = 4, 4               // 1, 2, 3, 4
	// -- a == 0 && b == l - 1
	// s, e = 0, 4               // nil
	// -- 0 < a < b < l
	// s, e = 1, 3               // 1, 5
	// a = 0
	// s, e = 0, 2               // 4, 5
	// b = l - 1
	// s, e = 2, 4               // 1, 2

	// Pssgl.RemoveList(s, e)

	// dd(Pssgl.Sync().Value)

	t.Log(FuncName + " ... ok!")
}

func Test_InsertList(t *testing.T) {
	FuncName := "GdLists.InsertList()"

	ssgl, iigl := getGdListIns()
	Pssgl := &ssgl
	Piigl := &iigl
	Pssgl.New()
	Piigl.New()

	// Pssgl.InsertList(2, Piigl)

	// var nssgl GdList
	// nssmi := GdInterface{[]string{"a", "b", "c", "d", "e"}}
	// nssgl.Value = nssmi.ToInterfaces()
	// Pnssgl := (&nssgl).New()

	// Pssgl.InsertList(0, Pnssgl)
	// Pssgl.InsertList(2, Pnssgl)
	// Pssgl.InsertList(4, Pnssgl)
	// Pssgl.InsertList(5, Pnssgl)

	// dd(Pssgl.Sync().Value)

	t.Log(FuncName + " ... ok!")
}

func Test_Insert(t *testing.T) {
	FuncName := "GdLists.Insert()"

	ssgl, _ := getGdListIns()
	Pssgl := &ssgl
	Pssgl.New()

	// var a interface{} = 50
	// Pssgl.Insert(2, a)

	// var b interface{} = "100"
	// Pssgl.Insert(2, b)

	// dd(Pssgl.Value)

	t.Log(FuncName + " ... ok!")
}
