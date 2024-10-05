package utils

func ZeroSliceMxN[T any](m, n int) [][]T {
	s := make([][]T, m)
	for ni := range m {
		s[ni] = make([]T, n)
	}
	return s
}
