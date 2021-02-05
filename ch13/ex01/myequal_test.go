package myEqual

import (
	"reflect"
	"testing"
)

func TestMyEqual(t *testing.T) {
	testInt := struct {
		x    int
		y    int
		want bool
	}{1, 1, true}

	testFloat := struct {
		x    float64
		y    float64
		want bool
	}{1.00000000000001, 1.00000000001, true}

	if MyEqual(reflect.ValueOf(testInt.x), reflect.ValueOf(testInt.y)) != testInt.want {
		t.Error("testInt error")
	}
	if MyEqual(reflect.ValueOf(testFloat.x), reflect.ValueOf(testFloat.y)) != testFloat.want {
		t.Error("testFloat error")
	}
}
