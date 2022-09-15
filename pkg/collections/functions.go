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

func Sort[T any](rows []T, greater func(a, b T) bool) []T {
	for i := 0; i < len(rows); i++ {
		for j := i + 1; j < len(rows); j++ {
			if greater(rows[i], rows[j]) {
				rows[i], rows[j] = rows[j], rows[i]
			}
		}
	}

	return rows
}
