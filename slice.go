package reflectio

import "reflect"

// MakeSlice will initialize a new slice
func MakeSlice(value interface{}, tagKey string) (s Slice) {
	rtype := reflect.TypeOf(value)
	return makeSlice(rtype, tagKey)
}

func makeSlice(rtype reflect.Type, tagKey string) (s Slice) {
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	numFields := rtype.NumField()
	s = make(Slice, numFields)
	for i := 0; i < numFields; i++ {
		field := rtype.Field(i)
		s[i] = makeField(i, field.Type.Kind())
	}

	return
}

// Slice represents fields of a reflected value
type Slice []Field
