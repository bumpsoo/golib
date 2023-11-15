package array

func Map[T, U any](arr []T, f func(T) U) []U {
	ret := make([]U, len(arr))
	for idx, val := range arr {
		ret[idx] = f(val)
	}
	return ret
}

func Filter[T any](arr []T, f func(T) bool) []T {
	ret := []T{}
	for _, val := range arr {
		if f(val) {
			ret = append(ret, val)
		}
	}
	return ret
}
