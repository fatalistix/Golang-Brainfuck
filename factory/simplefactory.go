package factory

import (
	"errors"
	"reflect"
)

type SimpleFactory struct {
	registry map[string]reflect.Type
}

func NewSimpleFactory() *SimpleFactory {
	return &SimpleFactory{make(map[string]reflect.Type)}
}

func (f *SimpleFactory) Register(command string, operation Operation) error {
	if f.registry[command] != nil {
		return errors.New("this command is already registered")
	}
	f.registry[command] = reflect.TypeOf(operation)
	return nil
}

func (f *SimpleFactory) CreateOperationInstance(command string) (Operation, error) {
	if f.registry[command] == nil {
		return nil, errors.New("operation '" + command + "' is not registered")
	}
	return reflect.New(f.registry[command]).Elem().Interface().(Operation), nil
}
