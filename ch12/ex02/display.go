package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Display(name string, x interface{}, cycle int) {
	fmt.Printf("Display %s (%T) with %d cycles:\n", name, x, cycle)
	display(name, reflect.ValueOf(x), 0, cycle)
}

func display(path string, v reflect.Value, start, cycle int) {
	if start == cycle {
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), start+1, cycle)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), start+1, cycle)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatMapKey(key)), v.MapIndex(key), start+1, cycle)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), start+1, cycle)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), start+1, cycle)
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func formatMapKey(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		keys := []string{}
		for i := 0; i < v.NumField(); i++ {
			keys = append(keys, formatAtom(v.Field(i)))
		}
		return "{" + strings.Join(keys, " ") + "}"
	case reflect.Array:
		keys := []string{}
		for i := 0; i < v.Len(); i++ {
			keys = append(keys, formatAtom(v.Index(i)))
		}
		return "[" + strings.Join(keys, " ") + "]"
	default:
		return formatAtom(v)

	}
}

type Cycle struct {
	Value int
	Tail  *Cycle
}

func main() {
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c, 100)
}
