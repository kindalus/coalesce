package coalesce

// Return the first NotNil element in the parameter list.
func NotNil[T any](values ...T) T {
	isNotNil := func(v T) bool { return interface{}(v) != nil }

	return Func(isNotNil, values...)
}

type EmptyConstraint interface {
	[]interface{} | ~string
}

// Return the first NotNil element in the parameter list.
func NotEmpty[T EmptyConstraint](values ...T) T {

	isNotEmpty := func(v T) bool { return len(v) > 0 }

	return Func(isNotEmpty, values...)
}

// Return the first element in the parameter list that is equal to the given target.
func Equal[T any](target T, values ...T) T {

	isEqual := func(value T) bool {
		return interface{}(value) == interface{}(target)
	}

	return Func(isEqual, values...)
}

// Return the first element in the parameter list that satisfies the given condition.
func Func[T any](test func(T) bool, values ...T) T {

	for _, v := range values {

		if test(v) {
			return v
		}

	}

	return values[0]
}
