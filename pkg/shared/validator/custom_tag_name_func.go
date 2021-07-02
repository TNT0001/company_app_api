package validator

import "reflect"

// FormTagName func
func FormTagName(fld reflect.StructField) string {
	name := fld.Tag.Get("form")
	if name != "" {
		return name
	}

	name = fld.Tag.Get("json")
	if name != "" {
		return name
	}

	return fld.Name
}
