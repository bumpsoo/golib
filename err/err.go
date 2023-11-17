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
