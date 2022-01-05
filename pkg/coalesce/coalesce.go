package coalesce

func Nil[T any](values ...T) T {
	isNil := func(v T) bool { return interface{}(v) == nil }

	return CoalesceFn(isNil, values...)
}

type EmptyConstraint interface {
	[]interface{} | ~string
}

func Empty[T EmptyConstraint](values ...T) T {

	isEmpty := func(v T) bool { return len(v) == 0 }

	return CoalesceFn(isEmpty, values...)
}

func Coalesce[T any](target interface{}, values ...T) T {

	isEqual := func(value T) bool {
		return interface{}(value) == target
	}

	return CoalesceFn(isEqual, values...)
}

func CoalesceFn[T any](testFn func(T) bool, values ...T) T {

	for _, v := range values {

		if !testFn(v) {
			return v
		}

	}

	return values[0]
}
