package structfield

import (
	"fmt"
	"reflect"
)

// Copy struct data from src to dst, when both
// structs have different fields. dst must be a
// pointer, otherwise no data can be copied.
func Copy(dst, src interface{}) error {
	typeDst := reflect.TypeOf(dst)
	if typeDst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not a pointer")
	}

	valDst := reflect.ValueOf(dst).Elem()
	valSrc := reflect.ValueOf(src)
	typeSrc := reflect.TypeOf(src)

	for i := 0; i < valSrc.NumField(); i++ {
		srcField := typeSrc.Field(i)
		dstField := valDst.FieldByName(srcField.Name)

		if !dstField.IsValid() {
			continue
		}

		if dstField.Kind() != valSrc.Field(i).Kind() {
			continue
		}

		srcTag := typeSrc.Field(i).Tag
		if srcTag.Get("structfield") == "nocopy" {
			continue
		}

		dstField.Set(valSrc.Field(i))
	}
	return nil
}
