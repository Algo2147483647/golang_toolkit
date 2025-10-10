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

// ToInt64 converts an interface{} to int64
func ToInt64(v interface{}) int64 {
	switch value := v.(type) {
	case int:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case int64:
		return value
	case uint:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint32:
		return int64(value)
	case uint64:
		return int64(value)
	case float32:
		return int64(value)
	case float64:
		return int64(value)
	case string:
		if i, err := strconv.ParseInt(value, 10, 64); err == nil {
			return i
		}
		return 0
	default:
		return 0
	}
}

// ToString converts an interface{} to string
func ToString(v interface{}) string {
	switch value := v.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		// For complex types, use JSON marshaling
		if data, err := json.Marshal(value); err == nil {
			return string(data)
		}
		return ""
	}
}

// ToIntSlice converts an interface{} to []int
func ToIntSlice(v interface{}) []int {
	switch value := v.(type) {
	case []int:
		return value
	case []interface{}:
		result := make([]int, len(value))
		for i, item := range value {
			result[i] = ToInt(item)
		}
		return result
	case []string:
		result := make([]int, len(value))
		for i, item := range value {
			result[i] = ToInt(item)
		}
		return result
	default:
		// Try to convert using JSON
		result, err := InterfaceToObject[[]int](v)
		if err == nil {
			return result
		}
		return []int{}
	}
}

// ToStringMapString converts an interface{} to map[string]string
func ToStringMapString(v interface{}) map[string]string {
	switch value := v.(type) {
	case map[string]string:
		return value
	case map[string]interface{}:
		result := make(map[string]string)
		for k, val := range value {
			result[k] = ToString(val)
		}
		return result
	case map[interface{}]interface{}:
		result := make(map[string]string)
		for k, val := range value {
			result[ToString(k)] = ToString(val)
		}
		return result
	default:
		// Try to convert using JSON
		result, err := InterfaceToObject[map[string]string](v)
		if err == nil {
			return result
		}
		return make(map[string]string)
	}
}

// ToStringMapInterface converts an interface{} to map[string]interface{}
func ToStringMapInterface(v interface{}) map[string]interface{} {
	switch value := v.(type) {
	case map[string]interface{}:
		return value
	case map[string]string:
		result := make(map[string]interface{})
		for k, val := range value {
			result[k] = val
		}
		return result
	case map[interface{}]interface{}:
		result := make(map[string]interface{})
		for k, val := range value {
			result[ToString(k)] = val
		}
		return result
	default:
		// Try to convert using JSON
		result, err := InterfaceToObject[map[string]interface{}](v)
		if err == nil {
			return result
		}
		return make(map[string]interface{})
	}
}
