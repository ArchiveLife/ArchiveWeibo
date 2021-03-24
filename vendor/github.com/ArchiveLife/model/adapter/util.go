package adapter

import "reflect"

func CreateInstanceByValues(t reflect.Type, values ...*OptionValue) reflect.Value {
	structureValue := reflect.Indirect(reflect.New(t))
	for _, value := range values {
		fieldValue := structureValue.FieldByName(value.Name)
		if fieldValue.IsValid() {
			if value.Value != nil {
				fieldValue.Set(reflect.ValueOf(value.Value))
			}
		}
	}
	return structureValue
}

func FillInstanceByValues(v interface{}, values ...*OptionValue) reflect.Value {
	structureValue := reflect.Indirect(reflect.ValueOf(v))
	for _, value := range values {
		fieldValue := structureValue.FieldByName(value.Name)
		if fieldValue.IsValid() {
			if value.Value != nil {
				fieldValue.Set(reflect.ValueOf(value.Value))
			}
		}
	}
	return structureValue
}
