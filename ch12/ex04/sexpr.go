package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, v reflect.Value, indent int, br bool) error {
	tabs := ""
	if br {
		tabs += "\n"
	}
	for i := 0; i < indent; i++ {
		tabs += "\t"
	}
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString(tabs + "nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%s%d", tabs, v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%s%d", tabs, v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%s%q", tabs, v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), indent+1, true)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteString(tabs)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i), indent+1, true); err != nil {
				return err
			}
		}
		buf.WriteString(tabs)
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteString(tabs)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+1, false); err != nil {
				return err
			}
			buf.WriteByte(')')
			buf.WriteString("\n")
		}
		buf.WriteString(tabs)
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteString("\n")
		buf.WriteString(tabs)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key, 0, false); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), 0, false); err != nil {
				return err
			}
			buf.WriteByte(')')
			buf.WriteString("\n")
		}
		buf.WriteString(tabs)
		buf.WriteByte(')')

	case reflect.Bool:
		buf.WriteString(tabs)
		if v.Bool() {
			fmt.Fprint(buf, "t")
		} else {
			fmt.Fprint(buf, "nil")
		}

	case reflect.Float32, reflect.Float64:
		buf.WriteString(tabs)
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128:
		buf.WriteString(tabs)
		fmt.Fprintf(buf, "#C(%g, %g)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		buf.WriteString(tabs)
		if v.IsNil() {
			fmt.Fprintf(buf, "nil")
		} else {
			var b bytes.Buffer
			encode(&b, v.Elem(), indent+1, false)
			fmt.Fprintf(buf, "(%q %s)", v.Elem().Type(), b.String())
		}

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0, false); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
