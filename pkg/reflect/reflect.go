package reflect

import (
	"reflect"
	"strconv"
	"strings"
)

func Reflect(re interface{}, data map[string]string) interface{} {
	t := reflect.TypeOf(re)
	v := reflect.ValueOf(re)
	n := t.Elem().NumField()

	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)
		key, ok := field.Tag.Lookup("json")
		if !ok {
			key = field.Name
		}
		value, ok := data[strings.ToLower(key)]
		if ok {
			// fill config
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				intValue, err := strconv.ParseUint(value, 10, 64)
				if err == nil {
					fieldVal.SetUint(intValue)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(intValue)
				}
			case reflect.Float64, reflect.Float32:
				intValue, err := strconv.ParseFloat(value, 64)
				if err == nil {
					fieldVal.SetFloat(intValue)
				}
			case reflect.Bool:
				boolValue := ("yes" == value || "true" == value)
				fieldVal.SetBool(boolValue)
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					slice := strings.Split(value, ",")
					fieldVal.Set(reflect.ValueOf(slice))
				}
			}
		}
	}
	return re
}
