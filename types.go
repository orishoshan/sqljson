package sqljson

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type JSON[T any] struct {
	Item T
}

func From[T any](input T) JSON[T] {
	return JSON[T]{Item: input}
}

func (j *JSON[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, &j.Item)
}

func (j JSON[T]) Value() (driver.Value, error) {
	itemValue := reflect.ValueOf(j.Item)
	switch itemValue.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		if itemValue.IsNil() {
			return nil, nil
		}
	}
	return json.Marshal(j.Item)
}

func (j JSON[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Item)
}

func (j *JSON[T]) UnmarshalJSON(data []byte) error {
	var out JSON[T]
	err := json.Unmarshal(data, &out.Item)
	*j = out
	return err
}
