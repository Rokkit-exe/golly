package utils

import (
	"fmt"
	"reflect"
)

func PrintStruct(v any) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	fmt.Println("{")
	for i := range val.NumField() {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf("    %s: %v\n", field.Name, value.Interface())
	}
	fmt.Println("}")
}
