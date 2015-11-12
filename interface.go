package godash

import (
	"reflect"
)

type GdInterface struct {
	Value interface{}
}

func (mi *GdInterface) TypeOf() string {
	return typeof(mi.Value)
}

func (mi *GdInterface) IsSlice() bool {
	return IsSlice(mi.Value)
}

func (mi *GdInterface) ToInterfaces() []interface{} {
	var r []interface{}
	if mi.IsSlice() { // 基于反射的迭代器
		value := reflect.ValueOf(mi.Value)
		for i := 0; i < value.Len(); i++ {
			v := value.Index(i).Interface()
			r = append(r, v)
		}
	} else {
		return nil
	}
	return r
}

type GdInterfaces struct {
	Value []interface{}
}

func (gis *GdInterfaces) ToStrings() []string {
	var r []string
	for _, v := range gis.Value {
    r = append(r, v.(string))
	}
	return r
}
