package common

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
)

// CompareJson compares two JSON strings and returns whether they are equal,
// a log of differences, and any error encountered.
func CompareJson(ctx context.Context, a, b string, path string, pathPassList []string) (bool, map[string]string, error) {
	var data1, data2 interface{}
	if err := json.Unmarshal([]byte(a), &data1); err != nil {
		return false, map[string]string{}, err
	}
	if err := json.Unmarshal([]byte(b), &data2); err != nil {
		return false, map[string]string{}, err
	}

	compareLog := map[string]string{}
	result := CompareInterfaces(ctx, data1, data2, path, pathPassList, compareLog)
	return result, compareLog, nil
}

// CompareInterfaces recursively compares two interfaces and logs differences.
func CompareInterfaces(ctx context.Context, data1, data2 interface{}, path string, pathPassList []string, compareLog map[string]string) bool {
	if Contains[string](pathPassList, path) {
		return true
	}

	a := reflect.ValueOf(data1)
	b := reflect.ValueOf(data2)
	areEqual := true

	switch {
	case !a.IsValid() && !b.IsValid():
		return true

	case !a.IsValid() || !b.IsValid():
		compareLog[path] += fmt.Sprintf("One of the values is invalid; ")
		return false

	case a.Kind() != b.Kind():
		if reflect.DeepEqual(a.Interface(), b.Interface()) {
			return true
		}
		compareLog[path] += fmt.Sprintf("Kind is not equal and DeepEqual is false: First = %v (%v), Second = %v (%v); ", a.Interface(), a.Type(), b.Interface(), b.Type())
		return false

	case a.Kind() == reflect.Ptr:
		if a.IsNil() && b.IsNil() {
			return true
		}

		if a.IsNil() || b.IsNil() {
			compareLog[path] += fmt.Sprintf("Pointer is nil: First = %v, Second = %v; ", a.Interface(), b.Interface())
			return false
		}

		return CompareInterfaces(ctx, a.Elem().Interface(), b.Elem().Interface(), path, pathPassList, compareLog)

	case a.Kind() == reflect.Struct:
		for i := 0; i < a.NumField(); i++ {
			field := a.Type().Field(i)
			// Skip unexported fields
			if !field.IsExported() {
				continue
			}
			nestedPath := path + "." + field.Name
			if !CompareInterfaces(ctx, a.Field(i).Interface(), b.Field(i).Interface(), nestedPath, pathPassList, compareLog) {
				areEqual = false
			}
		}

	case a.Kind() == reflect.Map:
		if a.Len() != b.Len() {
			compareLog[path] += fmt.Sprintf("Map lengths are different: First = %v, Second = %v; ", a.Len(), b.Len())
			areEqual = false
		}

		// Merge keys from both maps to form a complete set
		allKeys := make(map[interface{}]reflect.Value)
		for _, key := range a.MapKeys() {
			allKeys[key.Interface()] = key
		}
		for _, key := range b.MapKeys() {
			allKeys[key.Interface()] = key
		}

		for _, key := range allKeys {
			val1, ok1 := getMapValue(a, key)
			val2, ok2 := getMapValue(b, key)

			if !ok1 {
				compareLog[path] += fmt.Sprintf("First map error: missing key = %v; ", key.Interface())
				areEqual = false
			}

			if !ok2 {
				compareLog[path] += fmt.Sprintf("Second map error: missing key = %v; ", key.Interface())
				areEqual = false
			}

			if ok1 && ok2 {
				nestedPath := path + "." + key.String()
				if !CompareInterfaces(ctx, val1, val2, nestedPath, pathPassList, compareLog) {
					areEqual = false
				}
			}
		}

	case a.Kind() == reflect.Slice || a.Kind() == reflect.Array:
		if a.Len() != b.Len() {
			compareLog[path] += fmt.Sprintf("Slice/Array lengths are different: First = %v, Second = %v; ", a.Len(), b.Len())
			areEqual = false
		}

		maxLen := a.Len()
		if b.Len() > maxLen {
			maxLen = b.Len()
		}

		for i := 0; i < maxLen; i++ {
			var val1, val2 interface{}

			if i < a.Len() {
				val1 = a.Index(i).Interface()
			} else {
				compareLog[path] += fmt.Sprintf("Slice/Array error, index %v is out of the length of First %v; ", i, a.Len())
				areEqual = false
				continue
			}

			if i < b.Len() {
				val2 = b.Index(i).Interface()
			} else {
				compareLog[path] += fmt.Sprintf("Slice/Array error, index %v is out of the length of Second %v; ", i, b.Len())
				areEqual = false
				continue
			}

			nestedPath := fmt.Sprintf("%s[%d]", path, i)
			if !CompareInterfaces(ctx, val1, val2, nestedPath, pathPassList, compareLog) {
				areEqual = false
			}
		}

	case a.Kind() == reflect.String:
		str1 := a.Interface().(string)
		str2 := b.Interface().(string)

		var parsed1, parsed2 interface{}
		err1 := json.Unmarshal([]byte(str1), &parsed1)
		err2 := json.Unmarshal([]byte(str2), &parsed2)
		if err1 == nil && err2 == nil { // If both can be parsed as JSON, continue comparing the parsed structures
			return CompareInterfaces(ctx, parsed1, parsed2, path, pathPassList, compareLog)
		}

		if str1 != str2 { // Otherwise compare the raw strings
			compareLog[path] += fmt.Sprintf("String fields are different: First = %v, Second = %v; ", str1, str2)
			return false
		}
		return true

	default:
		if reflect.DeepEqual(a.Interface(), b.Interface()) {
			return true
		}
		compareLog[path] += fmt.Sprintf("Fields are different: First = %v, Second = %v; ", a.Interface(), b.Interface())
		return false
	}
	return areEqual
}
