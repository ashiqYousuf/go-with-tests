package treflection

import (
	"reflect"
)

// ! PROBLEM :- write a function walk(x interface{}, fn func(string))
// ! which takes a struct x and calls fn for all strings fields found inside.

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fn func(input string)) {
	// ? NumField returns the number of fields in the struct v.
	// ? It panics if v's Kind is not Struct.
	val := getValue(x)
	// walkValue := func(value reflect.Value) {
	// 	walk(value.Interface(), fn)
	// }

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
}
