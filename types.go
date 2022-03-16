package sqljson

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
