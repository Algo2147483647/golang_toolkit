package common

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
