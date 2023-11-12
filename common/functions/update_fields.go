package functions

import (
	"reflect"
	"time"
)

func UpdateFields(res interface{}, body interface{}) interface{} {
	now := time.Now().UTC()

	resValue := reflect.ValueOf(res).Elem()
	bodyValue := reflect.ValueOf(body).Elem()

	for i := 0; i < bodyValue.NumField(); i++ {
		bodyField := bodyValue.Field(i)
		resField := resValue.Field(i)

		if bodyField.Kind() == reflect.String && bodyField.String() != "" {
			resField.SetString(bodyField.String())
		}
	}

	updatedAtField := resValue.FieldByName("UpdatedAt")
	if updatedAtField.IsValid() && updatedAtField.CanSet() {
		updatedAtField.Set(reflect.ValueOf(&now))
	}

	return res
}
