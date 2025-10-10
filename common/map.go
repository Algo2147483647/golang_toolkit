package common

import "reflect"

// getMapValue retrieves the value associated with the given key in the map.
// It returns the value and a boolean indicating if the key exists.
func getMapValue(m reflect.Value, key interface{}) (interface{}, bool) {
	keyVal := reflect.ValueOf(key)
	if keyVal.Type().AssignableTo(m.Type().Key()) {
		val := m.MapIndex(keyVal)
		if val.IsValid() {
			return val.Interface(), true
		}
	}
	return nil, false
}
