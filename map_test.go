package reflectio

import (
	"log"
	"reflect"
	"strconv"
	"testing"
)

var (
	mapSink Map
)

func TestMap_SetValueAsString(t *testing.T) {
	var (
		test testStruct
		err  error
	)

	m := MakeMap(test, "reflectio")
	rval := reflect.ValueOf(&test)

	if err = m.SetValueAsString(rval, "int1", "13"); err != nil {
		t.Fatal(err)
	}

	if err = m.SetValueAsString(rval, "uint1", "1337"); err != nil {
		t.Fatal(err)
	}

	if err = m.SetValueAsString(rval, "float1", "3.14"); err != nil {
		t.Fatal(err)
	}

	if err = m.SetValueAsString(rval, "bool1", "true"); err != nil {
		t.Fatal(err)
	}

	if err = m.SetValueAsString(rval, "string1", "Hello world!"); err != nil {
		t.Fatal(err)
	}

	if test.Int1 != 13 {
		t.Fatalf("invalid value, expected %d and received %d", 13, test.Int1)
	}

	if test.Uint1 != 1337 {
		t.Fatalf("invalid value, expected %d and received %d", 1337, test.Uint1)
	}

	if test.Float1 != 3.14 {
		t.Fatalf("invalid value, expected %f and received %f", 3.14, test.Float1)
	}

	if test.Bool1 != true {
		t.Fatalf("invalid value, expected %v and received %v", true, test.Bool1)
	}

	if test.String1 != "Hello world!" {
		t.Fatalf("invalid value, expected \"%s\" and received \"%s\"", "Hello world!", test.String1)
	}
}

func TestMap_SetValueAsString_int8(t *testing.T) {
	type testStruct struct {
		Int1 int8 `reflectio:"int1"`
	}

	var (
		test testStruct
		err  error
	)

	m := MakeMap(test, "reflectio")
	rval := reflect.ValueOf(&test)

	if err = m.SetValueAsString(rval, "int1", "13"); err != nil {
		t.Fatal(err)
	}

	if test.Int1 != 13 {
		t.Fatalf("invalid value, expected %d and received %d", 13, test.Int1)
	}
}

func BenchmarkMakeMap(b *testing.B) {
	var test testStruct
	for i := 0; i < b.N; i++ {
		mapSink = MakeMap(test, "reflectio")
	}

	b.ReportAllocs()
}

func BenchmarkMap_SetValueAsString(b *testing.B) {
	var (
		test testStruct
		err  error
	)

	m := MakeMap(test, "reflectio")

	for i := 0; i < b.N; i++ {
		rval := reflect.ValueOf(&test)
		if err = m.SetValueAsString(rval, "int1", "13"); err != nil {
			b.Fatal(err)
		}

		if err = m.SetValueAsString(rval, "uint1", "1337"); err != nil {
			b.Fatal(err)
		}

		if err = m.SetValueAsString(rval, "float1", "3.14"); err != nil {
			b.Fatal(err)
		}

		if err = m.SetValueAsString(rval, "bool1", "true"); err != nil {
			b.Fatal(err)
		}

		if err = m.SetValueAsString(rval, "string1", "Hello world!"); err != nil {
			b.Fatal(err)
		}
	}

	b.ReportAllocs()
}

func BenchmarkRaw_SetValueAsString(b *testing.B) {
	var (
		test testStruct
		err  error
	)

	for i := 0; i < b.N; i++ {
		if test.Int1, err = strconv.ParseInt("13", 10, 64); err != nil {
			return
		}

		if test.Uint1, err = strconv.ParseUint("1337", 10, 64); err != nil {
			return
		}

		if test.Float1, err = strconv.ParseFloat("3.14", 64); err != nil {
			return
		}

		if test.Bool1, err = strconv.ParseBool("true"); err != nil {
			return
		}

		test.String1 = "Hello world!"
	}

	b.ReportAllocs()
}

type testStruct struct {
	Int1 int64 `reflectio:"int1"`
	Int2 int64 `reflectio:"int2"`
	Int3 int64 `reflectio:"int3"`
	Int4 int64 `reflectio:"int4"`

	Uint1 uint64 `reflectio:"uint1"`
	Uint2 uint64 `reflectio:"uint2"`
	Uint3 uint64 `reflectio:"uint3"`
	Uint4 uint64 `reflectio:"uint4"`

	Float1 float64 `reflectio:"float1"`
	Float2 float64 `reflectio:"float2"`
	Float3 float64 `reflectio:"float3"`
	Float4 float64 `reflectio:"float4"`

	String1 string `reflectio:"string1"`
	String2 string `reflectio:"string2"`
	String3 string `reflectio:"string3"`
	String4 string `reflectio:"string4"`

	Bool1 bool `reflectio:"bool1"`
	Bool2 bool `reflectio:"bool2"`
	Bool3 bool `reflectio:"bool3"`
	Bool4 bool `reflectio:"bool4"`
}

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
