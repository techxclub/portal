package domain

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"

	"github.com/techx/portal/utils"
)

var (
	_ JSONWrapperInterface = new(JSONWrapper[GoogleOAuthDetails])
	_ JSONWrapperInterface = new(JSONWrapper[TechnicalInformation])
	_ JSONWrapperInterface = new(JSONWrapper[MentorConfig])
)

type JSONWrapper[T any] struct {
	data T
}

type JSONWrapperInterface interface {
	encoding.TextUnmarshaler
	encoding.TextMarshaler
	sql.Scanner
	driver.Valuer
}

func (r *JSONWrapper[T]) GetData() T {
	return r.data
}

func (r *JSONWrapper[T]) SetData(d T) {
	r.data = d
}

func (r JSONWrapper[T]) Value() (driver.Value, error) {
	bytes, _ := r.encode()
	return bytes, nil
}

func (r *JSONWrapper[T]) Scan(src interface{}) error {
	ciphertext, err := utils.ScanBytes(src)
	if err != nil {
		return err
	}

	if err := r.decode(ciphertext); err != nil {
		return err
	}

	return nil
}

func (r *JSONWrapper[T]) UnmarshalText(b []byte) error {
	if err := r.decode(b); err != nil {
		return err
	}
	return nil
}

func (r JSONWrapper[T]) MarshalText() ([]byte, error) {
	encodedData, err := r.encode()
	if err != nil {
		return nil, err
	}

	return encodedData, nil
}

func (r JSONWrapper[T]) encode() ([]byte, error) {
	data, err := json.Marshal(r.data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *JSONWrapper[T]) decode(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, &r.data)
}
