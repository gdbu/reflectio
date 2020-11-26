# Reflectio
Reflectio is a helper library to aid with utilizing Golang Reflection.

## Usage
### Map.SetValueAsString
```go
func ExampleMap_SetValueAsString() {
	// Test struct, in the real world - this should be declared globally
	type testStruct struct {
		Int1 int8 `reflectio:"int1"`
	}

	var value testStruct
	// Make reflection map for testStruct using a tagging key of "reflectio"
	m := MakeMap(value, "reflectio")
	// Create a reflect.Value from our value
	rval := reflect.ValueOf(&value)

	// Set the "int1" field with the value of 13
	if err := m.SetValueAsString(rval, "int1", "13"); err != nil {
		log.Fatalf("error setting value for int1: %v", err)
	}
}

```