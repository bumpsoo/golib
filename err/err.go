package err

type (
	wrap[T any] struct {
		Value T
		Error error
	}
)

func Wrap[T any](value T, err error) wrap[T] {
	return wrap[T]{Value: value, Error: err}
}

func Do[T, U any](w *wrap[T], f func(T) (U, error)) wrap[U] {
	var ret U
	var er error
	if w.Error != nil {
		return wrap[U]{Value: ret, Error: w.Error}
	} else {
		ret, er = f(w.Value)
		return wrap[U]{Value: ret, Error: er}
	}
}
