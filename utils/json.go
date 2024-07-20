package utils

import (
	"encoding/json"
	"fmt"
)

// ScanJSON is a helper function to deduplicate code for scanning database
// column value of type JSONB into a concrete Go type.
func ScanJSON(data, value interface{}) error {
	b, err := ScanBytes(data)
	if err != nil {
		return err
	}

	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, value)
}

func ScanBytes(data interface{}) ([]byte, error) {
	switch data := data.(type) {
	case string:
		return []byte(data), nil
	case []byte:
		return data, nil
	case nil:
		return nil, nil
	}
	return nil, fmt.Errorf("incompatible type for %T", data)
}
