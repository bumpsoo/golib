package stack

type Stack[T any] []T

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (ret T, ok bool) {
	l := len(*s)
	if l == 0 {
		return
	} else {
		ret = (*s)[l-1]
		*s = (*s)[:l-1]
		return ret, true
	}
}
