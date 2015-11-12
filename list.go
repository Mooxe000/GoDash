package godash

import (
	"container/list"
	"sort"
	// "reflect"
)

type GdList struct {
	Value []interface{}
	list  *list.List
}

func (gl *GdList) New() *GdList {
	l := list.New() // TODO ml.list.init()
	for _, v := range gl.Value {
		l.PushBack(v)
	}
	gl.list = l
	return gl
}

// TODO params int TO uint avoid border check

// (Sync) TODO arithmetic wait optimize
func (gl *GdList) Sync() *GdList {
	var r []interface{}
	l := gl.list
	for e := l.Front(); e != nil; e = e.Next() {
		r = append(r, e.Value)
	}
	gl.Value = r
	return gl
}

// TODO isSynced()

func (gl *GdList) Len() int {
	l := gl.list
	return l.Len()
}

func (gl *GdList) At(i int) *list.Element {
	l := gl.list
	ll := gl.Len()
	var r *list.Element
	if i < 0 || i >= ll {
		return nil
	} else {
		start := l.Front()
		switch i {
		case 0:
			r = start
		case ll - 1:
			r = l.Back()
		default:
			t := 0
			for e := start; e != nil; e = e.Next() {
				if t == i {
					r = e
					break
				} else {
					t++
				}
			}
		}
	}
	return r
}

func (gl *GdList) Getter(i int) interface{} {
	return gl.At(i).Value
}

func (gl *GdList) Setter(i int, v interface{}) *GdList {
	gl.At(i).Value = v
	return gl
}

// Remove
func (gl *GdList) Remove(i int) *GdList {
	if i < 0 || i >= gl.Len() {
		return gl
	}
	l := gl.list
	e := gl.At(i)
	l.Remove(e)
	return gl
}

func iCheck(i int, l int) int {
	var r int
	if i < 0 {
		r = 0
	} else if i >= l {
		r = l - 1
	} else {
		r = i
	}
	return r
}

func seCheck(s int, e int, l int) (int, int) {
	ii := []int{s, e}
	sort.Ints(ii)
	// dd(ii)
	a, b := ii[0], ii[1]
	if a < 0 && b < 0 {
		a, b = 0, 0
	} else if a >= l && b >= l {
		a, b = l-1, l-1
	} else if a < 0 && b >= l {
		a = 0
		b = l - 1
	} else if a < 0 && b >= 0 {
		a = 0
	} else if a >= 0 && b >= l {
		b = l - 1
	}
	return a, b
}

// Slice
func (gl *GdList) Slice(s int, e int) *GdList {
	l := gl.Len()
	a, b := seCheck(s, e, l)

	var ngl GdList
	v := gl.Value

	if a == b {
		var is []interface{}
		is = append(is, v[s])
		ngl.Value = is

	} else {
		ngl.Value = v[a : b+1]
	}

	Pgl := &ngl
	Pgl.New()
	return Pgl
}

func (gl *GdList) Chunk(ii ...int) []GdList {
	lii := len(ii)
	l := gl.Len()

	var gls []GdList

	if lii <= 0 {
		gls = append(gls, *gl)
		return gls
	}

	f := ii[0]
	if f < 0 || f >= l {

		gls = append(gls, *gl)
		return gls

	} else {

		s, e := 0, f
		// dd(s)
		// dd(e)

		for i := 1; ; i++ {
			if s != e {
				var ngl GdList
				ngl.Value = gl.Slice(s, e-1).Value
				ngl.New()
				gls = append(gls, ngl)
			}

			// start point
			if e >= l {
				break
			} else {
				s = e
			}

			// end point
			// -- step n
			var n int
			if lii == 1 { // one repeat

				if f == 0 {

					gls = append(gls, *gl)
					return gls

				} else {
					n = f
				}
			} else { // one by one
				if i < lii { // item left
					if ii[i] == 0 {
						continue
					} else {
						n = ii[i]
					}
				} else { // item empty
					n = l - s
				}
			}

			if s+n >= l {
				n = l - s
			}

			// prf("s: %d\n", s)
			// prf("e: %d\n", e)
			// prf("n: %d\n", n)

			e = s + n
		}
	}
	return gls
}

// Concat
func (gl *GdList) Concat(gls []GdList) *GdList {
	var ngl GdList
	Pngl := &ngl
	for _, v := range gls {
		if Pngl.list == nil {
			Pngl.list = v.list
		} else {
			Pngl.list.PushBackList(v.list)
		}
	}
	gl.list = Pngl.list
	return gl
}

// RemoveList
func (gl *GdList) RemoveList(s int, e int) *GdList {
	l := gl.Len()
	a, b := seCheck(s, e, l)

	if a == b {
		return gl.Remove(s)

	} else if a == 0 && b == l-1 {
		gl.list.Init()
		return gl

	} else if a == 0 && b < l-1 {
		gl.list = gl.Slice(b+1, l-1).list

	} else if a > 0 && b == l-1 {
		gl.list = gl.Slice(0, a-1).list

	} else { // a > 0 && b < l - 1

		var gls []GdList

		gla := gl.Slice(0, a-1)
		glb := gl.Slice(b+1, l-1)

		// dd(mla.Value)
		// dd(mlb.Value)

		gls = append(gls, *gla)
		gls = append(gls, *glb)

		gl.list = gl.Concat(gls).list
	}

	return gl
}

// InsertList
func (gl *GdList) InsertList(i int, ngl *GdList) *GdList {
	l := gl.Len()
	i = iCheck(i, l)

	Tgl := typeof(gl.Value[0])
	Tngl := typeof(ngl.Value[0])
	if Tgl != Tngl {
		return gl
	}

	var gls []GdList
	if i == 0 {

		gls = append(gls, *ngl)
		gls = append(gls, *gl)
		gl.Concat(gls)

	} else if i == l-1 {

		gls = append(gls, *gl)
		gls = append(gls, *ngl)
		gl.Concat(gls)

	} else {

		a := gl.Slice(0, i-1)
		b := gl.Slice(i, l-1)
		gls = append(gls, *a)
		gls = append(gls, *ngl)
		gls = append(gls, *b)
		gl.Concat(gls)

	}
	return gl
}

// Insert
func (gl *GdList) Insert(i int, ii interface{}) *GdList {
	l := gl.Len()
	i = iCheck(i, l)

	Tgl := typeof(gl.Value[0])
	Tii := typeof(ii)
	if Tgl != Tii {
		return gl
	}

	var v []interface{}
	v = append(v, ii)
	var ngl GdList
	ngl.Value = v
	Pngl := &ngl
	Pngl.New()

	gl.InsertList(i, Pngl).Sync()

	return gl
}

// Pop -- remove one at the back of list
func (gl *GdList) Pop() (interface{}, *GdList) {
	l := gl.Len()
	if l <= 1 {
		return nil, nil
	}
	p := gl.Getter(l - 1)
	gl.Remove(l - 1)
	// gl.Sync()
	return p, gl
}

// Push -- add one at the back of list
func (gl *GdList) Push(i interface{}) *GdList {
	l := gl.list
	l.PushBack(i)
	return gl
}

// Shift -- remove one at the front of list
func (gl *GdList) Shift() (interface{}, *GdList) {
	l := gl.Len()
	if l <= 1 {
		return nil, nil
	}
	p := gl.Getter(0)
	gl.Remove(0)
	// gl.Sync()
	return p, gl
}

// Unshift -- add one at the front of list
func (gl *GdList) Unshift(i interface{}) *GdList {
	l := gl.list
	l.PushFront(i)
	return gl
}

// (Without)
// (IndexOf)
// (Flatten)
// (Sort)
// (Compact) -- Maybe can add param type as "string"/"int"/"interface{}"

////////////////////////////////////////////////////
////////////////////////////////////////////////////
////////////////////////////////////////////////////
// func (ml *GdList) RemoveListReflect(s int, e int) *GdList {
//   if s == e {
//     return ml
//   }
//   if s > e {
//     t := e
//     s = t
//     e = s
//   }
//   l := ml.Len()
//   if s <= 0 {
//     ml.list = ml.Slice(e, l).list
//   } else if e >= l {
//     ml.list = ml.Slice(0, s).list
//   } else {
//     hb := reflect.ValueOf(ml.At(s-1)).Elem()
//     ff := reflect.ValueOf(ml.At(e+1)).Elem()
//
//     // 0 - Next 1 - Prev
//     // 2 - List 3 - Value
//     // dd(hb.Type().Field(n).Name)
//
//     // dd(hb.Field(1).CanSet())
//
//     // dd(hb.Field(3).Interface())
//     // var v interface{} = "-------------"
//     // hb.Field(3).Set(reflect.ValueOf(v))
//     // ff.Field(3).Set(reflect.ValueOf(v))
//
//     hbn := hb.Field(0)
//     ffp := ff.Field(1)
//
//     dd(hbn.CanSet())
//     dd(ffp.CanSet())
//
//     Phb := ml.At(s).Prev()
//     Pff := ml.At(e).Next()
//
//     pln(reflect.ValueOf(Phb).Type())
//     pln(reflect.ValueOf(Pff).Type())
//
//     pln(hbn.Type())
//     pln(ffp.Type())
//
//     // hbn.Set(reflect.ValueOf(&Phb))
//     // ffp.Set(reflect.ValueOf(&Pff))
//   }
//
//   return ml
// }
