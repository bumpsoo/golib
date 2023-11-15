package hash

func Map[C comparable, T, U any](m map[C]T, f func(T) U) map[C]U {
	ret := make(map[C]U, len(m))
	for key, val := range m {
		ret[key] = f(val)
	}
	return ret
}
