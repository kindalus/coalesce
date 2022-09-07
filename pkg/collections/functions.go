package collections

func Map[T, M any](rows []T, f func(r T) M) []M {
	result := make([]M, 0)

	for _, row := range rows {
		result = append(result, f(row))
	}

	return result
}

func Filter[T any](rows []T, f func(r T) bool) []T {
	result := make([]T, 0)

	for _, row := range rows {
		if f(row) {
			result = append(result, row)
		}
	}

	return result
}

func Any[T any](rows []T, f func(r T) bool) bool {
	for _, row := range rows {
		if f(row) {
			return true
		}
	}

	return false
}
