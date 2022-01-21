package coalesce

func Nil[T any](values ...T) T {
	isNil := func(v T) bool { return interface{}(v) != nil }

	return Func(isNil, values...)
}

type EmptyConstraint interface {
	[]interface{} | ~string
}

func Empty[T EmptyConstraint](values ...T) T {

	isEmpty := func(v T) bool { return len(v) > 0 }

	return Func(isEmpty, values...)
}

func Equal[T any](target T, values ...T) T {

	isEqual := func(value T) bool {
		return interface{}(value) == interface{}(target)
	}

	return Func(isEqual, values...)
}

func Func[T any](test func(T) bool, values ...T) T {

	for _, v := range values {

		if test(v) {
			return v
		}

	}

	return values[0]
}
