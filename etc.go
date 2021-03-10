package njson

import (
	"reflect"
	"strings"
)

func validTag(field reflect.StructField) bool {
	if field.Anonymous {
		return true
	}

	return !(field.Tag.Get(tag) == "" ||
		field.Tag.Get(tag) == "-")
}

func isStructureType(typ string) (ok bool) {
	switch typ {
	case reflect.Slice.String():
		ok = true
	case reflect.Map.String():
		ok = true
	case reflect.Struct.String():
		ok = true
	default:
		ok = false
	}

	if strings.Contains(typ, "[]") {
		ok = true
	}

	return
}
