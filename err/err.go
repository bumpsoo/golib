package err

type (
	wrap[T any] struct {
		Value T
		Error error
	}

	monad[T any] func() (T, error)
)

func Wrap[T any](value T, err error) monad[T] {
	return func() (T, error) {
		return value, err
	}
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

func Bind[T, U any](m monad[T], f func(T) (U, error)) monad[U] {
	return func() (U, error) {
		v, err := m()
		var ret U
		if err != nil {
			return ret, err
		}
		return f(v)
	}
}
