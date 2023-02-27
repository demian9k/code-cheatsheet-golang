package util

// Set Difference: A - B
func Difference[T comparable](a, b []T) []T {
	m := make(map[T]bool)

	for _, item := range b {
		m[item] = true
	}

	var diff []T
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
