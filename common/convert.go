package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// InterfaceToObject converts an interface{} to a specified type T.
// It first tries a direct type assertion, and if that fails, it uses JSON marshaling/unmarshaling.
func InterfaceToObject[T any](data interface{}) (result T, err error) {
	if data == nil {
		return result, fmt.Errorf("data is nil")
	}

	if value, ok := data.(T); ok {
		return value, nil
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, fmt.Errorf("failed to marshal data: %w", err)
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal data to type %T: %w", result, err)
	}

	return result, nil
}

// ToInt converts an interface{} to int
func ToInt(v interface{}) int {
	switch value := v.(type) {
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		return int(value)
	case uint64:
		return int(value)
	case float32:
		return int(value)
	case float64:
		return int(value)
	case string:
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
		return 0
	default:
		return 0
	}
}
