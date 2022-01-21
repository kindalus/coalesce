package coalesce

func NotNil[T any](values ...T) T {
	isNotNil := func(v T) bool { return interface{}(v) != nil }

	return Func(isNotNil, values...)
}

type EmptyConstraint interface {
	[]interface{} | ~string
}

func NotEmpty[T EmptyConstraint](values ...T) T {

	isNotEmpty := func(v T) bool { return len(v) > 0 }

	return Func(isNotEmpty, values...)
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
