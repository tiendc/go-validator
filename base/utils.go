package base

func ToMap[T comparable](s []T) map[T]struct{} {
	result := make(map[T]struct{}, len(s))
	for _, v := range s {
		result[v] = struct{}{}
	}
	return result
}

func IsIn[T comparable](s []T, list []T) int {
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

func IsNotIn[T comparable](s []T, list []T) int {
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

func IsUnique[T comparable](s []T) int {
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
