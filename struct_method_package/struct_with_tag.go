package struct_method_package

import (
	"fmt"
	"reflect"
)

type tagStruct struct {
	field1 string "this is field1"
	file2  string "this is field2"
}

func getTag() {
	tagStruct := tagStruct{field1: "1", file2: "2"}
	for i := 0; i < 2; i++ {
		ref := reflect.TypeOf(tagStruct)
		fieldStruct := ref.Field(i)
		fmt.Printf("field struct name is %s tag is %s\n", fieldStruct.Name, fieldStruct.Tag)
	}
}
