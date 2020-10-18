package runtimestruct

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"unicode"
	"unsafe"
)

func NewFromJSON(r io.Reader) (interface{}, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return generateStruct(m), nil
}

func generateStruct(m map[string]interface{}) interface{} {
	fields := make([]reflect.StructField, 0, len(m))
	values := make([]interface{}, 0, len(m))
	for name, value := range m {
		if object, isObject := value.(map[string]interface{}); isObject {
			value = generateStruct(object)
		}

		typ := reflect.TypeOf(value)
		if typ == nil {
			typ = reflect.TypeOf((*struct{})(nil)) // typed nil
		}
		path := ""
		if !isExported(name) {
			path = "github.com/GreenHedgehog/runtimestruct"
		}
		fields = append(fields, reflect.StructField{
			Name:    name,
			PkgPath: path,
			Type:    typ,
		})
		values = append(values, value)
	}

	object := reflect.New(reflect.StructOf(fields))
	for i := 0; i < len(m); i++ {
		field := object.Elem().Field(i)
		makeSettable(&field)
		if values[i] != nil {
			field.Set(reflect.ValueOf(values[i]))
		}
	}
	return object.Elem().Interface()
}

func makeSettable(field *reflect.Value) {
	if field.CanSet() {
		return
	}

	addr := field.UnsafeAddr()
	*field = reflect.NewAt(
		field.Type(),
		unsafe.Pointer(addr),
	).Elem()
}

func isExported(fieldName string) bool {
	for _, r := range fieldName {
		return unicode.IsUpper(r) && unicode.IsLetter(r)
	}
	// would never happen
	return false
}
