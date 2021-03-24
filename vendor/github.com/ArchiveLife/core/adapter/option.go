package adapter

import "reflect"

type Option struct {
	Order int
	// tech name for field
	Name string
	// readable label for ui
	Label *string
	// readable description for ui
	Description *string
	// user can ignore the value
	Optional  bool
	ValueType reflect.Kind
}

type OptionValue struct {
	Option
	Value interface{}
}
