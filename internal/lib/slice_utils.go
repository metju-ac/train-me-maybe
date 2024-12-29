package lib

import "slices"

func MapFunc[T any, U any](slice []T, mapFunc func(T) U) []U {
	mappedSlice := make([]U, 0, len(slice))
	for _, value := range slice {
		mappedSlice = append(mappedSlice, mapFunc(value))
	}

	return mappedSlice
}

func FilterFunc[T any](slice []T, filterFunc func(T) bool) []T {
	filteredSlice := make([]T, 0, len(slice))
	for _, value := range slice {
		if filterFunc(value) {
			filteredSlice = append(filteredSlice, value)
		}
	}

	return slices.Clip(filteredSlice)
}

func ToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool, len(slice))

	for i := 0; i < len(slice); i++ {
		set[slice[i]] = true
	}

	return set
}