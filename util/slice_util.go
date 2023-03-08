package util

// Generic function to remove duplicates from a slice of any type
func RemoveDuplicates2(slice interface{}) interface{} {
	m := make(map[interface{}]bool)
	result := make([]interface{}, 0)

	// Loop through the slice and add each element to the map
	for _, elem := range slice.([]interface{}) {
		if _, ok := m[elem]; !ok {
			m[elem] = true
			result = append(result, elem)
		}
	}

	return result
}

func RemoveDuplicates[T comparable](slice []T) []T {
	m := make(map[T]bool)
	result := make([]T, 0)

	for _, elem := range slice {
		if _, ok := m[elem]; !ok {
			m[elem] = true
			result = append(result, elem)
		}
	}

	return result
}
