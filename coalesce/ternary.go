package coalesce

func Ternary[T any](test bool, a T, b T) T {
	if test {
		return a
	}

	return b
}
