package utils

// Filter returns the slice filter f(s[i]),
// or -1 if none do.
func Filter[S ~[]E, E any](s S, f func(E) bool) S {
	var filteredSlice S

	for i := range s {
		if f(s[i]) {
			filteredSlice = append(filteredSlice, s[i])
		}
	}

	return filteredSlice
}
