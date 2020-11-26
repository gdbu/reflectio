package reflectio

import (
	"reflect"
	"strconv"
)

func setInt(target reflect.Value, value string, bitSize int) (err error) {
	var i64 int64
	if i64, err = strconv.ParseInt(value, 10, bitSize); err != nil {
		return
	}

	target.SetInt(i64)
	return
}

func setUint(target reflect.Value, value string, bitSize int) (err error) {
	var u64 uint64
	if u64, err = strconv.ParseUint(value, 10, bitSize); err != nil {
		return
	}

	target.SetUint(u64)
	return
}

func setFloat(target reflect.Value, value string, bitSize int) (err error) {
	var f64 float64
	if f64, err = strconv.ParseFloat(value, bitSize); err != nil {
		return
	}

	target.SetFloat(f64)
	return
}

func setBool(target reflect.Value, value string) (err error) {
	var b bool
	if b, err = strconv.ParseBool(value); err != nil {
		return
	}

	target.SetBool(b)
	return
}
