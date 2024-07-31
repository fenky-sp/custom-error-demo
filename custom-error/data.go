package customerror

import (
	"reflect"
	"strings"
)

var (
	allowDataAttachment = true

	// map of field name marked as PII
	// key must be in lowercase
	piiFieldNameMap = map[string]bool{
		"phoneno": true,
	}
)

type dataIterator struct {
	ContainInterface bool
}

func (di *dataIterator) iterateData(
	value reflect.Value,
	tag reflect.StructTag,
	fieldName string,
	fn func(reflect.Value, reflect.Kind, string, reflect.StructTag),
) {
	if di == nil {
		return
	}

	value = reflect.Indirect(value)

	kind := value.Kind()

	switch kind {

	case reflect.Interface:
		di.ContainInterface = true

	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			di.iterateData(value.Field(i), t.Field(i).Tag, t.Field(i).Name, fn)
			if di.ContainInterface {
				break
			}
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			di.iterateData(value.Index(i), tag, fieldName, fn)
			if di.ContainInterface {
				break
			}
		}

	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()

			tmp := reflect.New(v.Type()) // create zero value of same type as v
			tmp.Elem().Set(v)
			tmp = reflect.Indirect(tmp)

			di.iterateData(tmp, tag, fieldName, fn)

			value.SetMapIndex(k, tmp)
		}

	case reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fn(value, kind, fieldName, tag)

	}
}

func processData(value reflect.Value, kind reflect.Kind, fieldName string, tag reflect.StructTag) {
	tagValue := tag.Get(CtxErrTagKey)
	tagValueMap := make(map[string]string)
	if tagValue != "" {
		tagValueArr := strings.Split(tagValue, ",")
		for _, tagValue := range tagValueArr {
			keyValueArr := strings.Split(tagValue, "=")
			var (
				pKey   string = tagValue
				pValue string
			)
			if len(keyValueArr) == 2 {
				pKey = keyValueArr[0]
				pValue = keyValueArr[1]
			}
			tagValueMap[pKey] = pValue
		}
	}

	if isPII(value, fieldName, tagValueMap) {
		maskData(value)
	}
}

func isPII(
	value reflect.Value,
	fieldName string,
	tagValueMap map[string]string,
) bool {
	// check whitelisted field name
	if _, exists := piiFieldNameMap[strings.ToLower(fieldName)]; exists {
		return true
	}

	// check tag
	if _, exists := tagValueMap[CtxErrTagValuePii]; exists {
		return true
	}

	// TODO check value with regex

	return false
}

func maskData(value reflect.Value) {
	if !value.CanSet() {
		// unexported fields cannot be set
		return
	}
	switch value.Kind() {
	case reflect.String:
		maskString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		maskInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		maskUint(value)
	}
}

func maskString(value reflect.Value) {
	value.SetString("***") // TODO
}

func maskInt(value reflect.Value) {
	value.SetInt(888) // TODO
}

func maskUint(value reflect.Value) {
	value.SetUint(888) // TODO
}
