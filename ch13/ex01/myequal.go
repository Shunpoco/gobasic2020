package myEqual

import (
	"math"
	"reflect"
)

func MyEqual(x, y reflect.Value) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()
	case reflect.Float32, reflect.Float64:
		return math.Abs(x.Float()-y.Float()) < 0.00000001
	default:
		return false
	}
}
