package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// reflect.ValueOf get the value of a given variable (x)
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}