package godash

import (
	"testing"
)

func Test_TypeOf(t *testing.T) {
	if r := typeof("Hello World!!!"); r != "string" {
		t.Error("typeof() ... failed!")
	} else {
		t.Log("typeof() ... ok!")
	}
}

// func (i Integer) ToString() string {
//   return strconv.Itoa(i)
// }

func Test_IsType(t *testing.T) {

	// s := "Hello World!!!"

	// dd(IsType(s, "int"))
	// dd(IsType(s, "string"))
	//
	// dd(IsInt(s))
	// dd(IsString(s))

	t.Log("typeof() ... ok!")
}
