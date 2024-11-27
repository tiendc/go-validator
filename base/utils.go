package base

// ToMap transforms a slice to a map with slice items as map keys
func ToMap[T comparable, S ~[]T](s S) map[T]struct{} {
	result := make(map[T]struct{}, len(s))
	for _, v := range s {
		result[v] = struct{}{}
	}
	return result
}

// IsIn returns -1 if every item of a slice is in another slice.
// Returns index of the first item if it is not in the target slice.
func IsIn[T comparable, S1 ~[]T, S2 ~[]T](s S1, list S2) int {
	if len(s) == 0 {
		return -1
	}
	if len(list) == 0 {
		return 0
	}
	m := ToMap(list)
	for i, v := range s {
		if _, ok := m[v]; !ok {
			return i
		}
	}
	return -1
}

// IsNotIn returns -1 if any item of a slice is not in another slice.
// Returns index of the first item if it is in the target slice.
func IsNotIn[T comparable, S1 ~[]T, S2 ~[]T](s S1, list S2) int {
	if len(s) == 0 || len(list) == 0 {
		return -1
	}
	m := ToMap(list)
	for i, v := range s {
		if _, ok := m[v]; ok {
			return i
		}
	}
	return -1
}

// IsUnique returns -1 if every item of a slice is unique.
// Returns index of the first item if it is a duplication of another.
func IsUnique[T comparable, S ~[]T](s S) int {
	length := len(s)
	if length <= 1 {
		return -1
	}
	seen := make(map[T]struct{}, length)
	for i := 0; i < length; i++ {
		v := s[i]
		if _, ok := seen[v]; ok {
			return i
		}
		seen[v] = struct{}{}
	}
	return -1
}

// IsUniqueBy returns -1 if every value returned by the key function is unique.
// Returns index of the first item if it is a duplication of another.
func IsUniqueBy[T any, U comparable, S ~[]T](s S, keyFn func(T) U) int {
	length := len(s)
	if length <= 1 {
		return -1
	}
	seen := make(map[U]struct{}, length)
	for i := 0; i < length; i++ {
		v := keyFn(s[i])
		if _, ok := seen[v]; ok {
			return i
		}
		seen[v] = struct{}{}
	}
	return -1
}
