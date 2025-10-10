package common

import (
	"encoding/json"
	"fmt"
)

// Contains checks if an element is in a slice, works with any type T
// Returns true if the element is found, false otherwise.
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

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
