package godash

import (
	"reflect"
	"strconv"
)

// TODO extends any type
func TypeOf(i interface{}) string {
	return reflect.TypeOf(i).String()
}

// TODO https://golang.org/pkg/reflect/#Kind
func IsType(i interface{}, ts string) bool {
	switch reflect.TypeOf(i).Kind() {

	// Interface
	case reflect.Interface:
		if ts == "slice" {
			return true
		} else {
			return false
		}

	// Bool
	case reflect.Bool:
		if ts == "bool" {
			return true
		} else {
			return false
		}
	// Int
	case reflect.Int:
		if ts == "int" {
			return true
		} else {
			return false
		}
	// String
	case reflect.String:
		if ts == "string" {
			return true
		} else {
			return false
		}

	// Struct
	case reflect.Struct:
		if ts == "struct" {
			return true
		} else {
			return false
		}
		// Func
	case reflect.Func:
		if ts == "func" {
			return true
		} else {
			return false
		}
	// Ptr
	case reflect.Ptr:
		if ts == "ptr" {
			return true
		} else {
			return false
		}

	// Array
	case reflect.Array:
		if ts == "array" {
			return true
		} else {
			return false
		}
	// Slice
	case reflect.Slice:
		if ts == "slice" {
			return true
		} else {
			return false
		}
	// Map
	case reflect.Map:
		if ts == "map" {
			return true
		} else {
			return false
		}
	// Chan
	case reflect.Chan:
		if ts == "chan" {
			return true
		} else {
			return false
		}

	// Default
	default:
		return false
	}
}

// Interface
func IsInterface(i interface{}) bool {
	return IsType(i, "interface")
}

// Bool
func IsBool(i interface{}) bool {
	return IsType(i, "bool")
}

// Int
func IsInt(i interface{}) bool {
	return IsType(i, "int")
}

// String
func IsString(i interface{}) bool {
	return IsType(i, "string")
}

// Struct
func IsStruct(i interface{}) bool {
	return IsType(i, "struct")
}

// Func
func IsFunc(i interface{}) bool {
	return IsType(i, "func")
}

// Ptr
func IsPtr(i interface{}) bool {
	return IsType(i, "prt")
}

// Array
func IsArray(i interface{}) bool {
	return IsType(i, "array")
}

// Slice
func IsSlice(i interface{}) bool {
	return IsType(i, "slice")
}

// Map
func IsMap(i interface{}) bool {
	return IsType(i, "map")
}

// Chan
func IsChan(i interface{}) bool {
	return IsType(i, "chan")
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

// without
// type Integer int
//
// func (a Integer) Less(b Integer) bool {
//     return a < b
// }
