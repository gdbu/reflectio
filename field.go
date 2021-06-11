package reflectio

import (
	"fmt"
	"reflect"
)

func makeField(fieldIndex int, kind reflect.Kind) (f Field) {
	f.fieldIndex = fieldIndex
	f.kind = kind
	return
}

// Field represents a reflected Field
type Field struct {
	// Field index for parent
	fieldIndex int
	// Kind of value
	kind reflect.Kind
}

func (f *Field) SetValueAsString(target reflect.Value, value string) (err error) {
	switch f.kind {
	case reflect.String:
		target.SetString(value)
		return
	case reflect.Int8:
		return setInt(target, value, 8)
	case reflect.Int16:
		return setInt(target, value, 16)
	case reflect.Int32:
		return setInt(target, value, 32)
	case reflect.Int64:
		return setInt(target, value, 64)
	case reflect.Uint8:
		return setUint(target, value, 8)
	case reflect.Uint16:
		return setUint(target, value, 16)
	case reflect.Uint32:
		return setUint(target, value, 32)
	case reflect.Uint64:
		return setUint(target, value, 64)
	case reflect.Float32:
		return setFloat(target, value, 64)
	case reflect.Float64:
		return setFloat(target, value, 64)
	case reflect.Bool:
		return setBool(target, value)

	default:
		err = fmt.Errorf("unsupported type provided, %s is not currently supported", f.kind)
		return
	}
}
