package api

import (
	"reflect"
)

func trimNil(val interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	varType := reflect.TypeOf(val)

	value := reflect.ValueOf(val)
	for i := 0; i < varType.NumField(); i++ {
		if !value.Field(i).CanInterface() ||
			value.Field(i).IsNil() {
			continue
		}

		tag, ok := varType.Field(i).Tag.Lookup("json")

		var fieldName string
		if ok && len(tag) > 0 {
			fieldName = tag
		} else {
			fieldName = varType.Field(i).Name
		}

		data[fieldName] = value.Field(i).Interface()
	}

	return data
}
