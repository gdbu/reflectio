package reflectio

import "reflect"

// MakeMap will initialize a new map
func MakeMap(value interface{}, tagKey string) (m Map) {
	rtype := reflect.TypeOf(value)
	return makeMap(rtype, tagKey)
}

func makeMap(rtype reflect.Type, tagKey string) (m Map) {
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}

	numFields := rtype.NumField()
	m = make(Map, numFields)
	for i := 0; i < numFields; i++ {
		field := rtype.Field(i)
		fieldValue, ok := field.Tag.Lookup(tagKey)
		if !ok {
			continue
		}

		m[fieldValue] = makeField(i, field.Type.Kind())
	}

	return
}

// Map represents fields of a reflected value
type Map map[string]Field

// SetValueAsString will attempt to set the string value of a given key within a provided target
// Note: The value will be attempted to converted to the appropriate type of the target
func (m Map) SetValueAsString(target reflect.Value, key, value string) (err error) {
	if target.Kind() == reflect.Ptr {
		target = target.Elem()
	}

	entry, ok := m[key]
	if !ok {
		return
	}

	field := target.Field(entry.fieldIndex)
	if setter, ok := field.Interface().(Setter); ok {
		return setter.SetValueAsString(value)
	}

	return entry.setValueAsString(field, value)
}
