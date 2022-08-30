package slices

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Filter performs the callback func on each entry, returning each entry where callback returns true.
func Filter[T any](slice []T, callback func(index int, value T) bool) []T {
	if len(slice) == 0 {
		return slice
	}
	var result []T
	for index, value := range slice {
		if callback(index, value) {
			result = append(result, value)
		}
	}
	return result
}

// Every returns true if callback evaluates to true for every entry in the slice.
func Every[T any](slice []T, callback func(index int, value T) bool) bool {
	if len(slice) == 0 {
		return false
	}
	for index, value := range slice {
		if !callback(index, value) {
			return false
		}
	}
	return true
}

// Some returns true if callback evaluates to true for any entry in the slice.
func Some[T any](slice []T, callback func(index int, value T) bool) bool {
	for index, value := range slice {
		if callback(index, value) {
			return true
		}
	}
	return false
}

// RemoveAt removes a value from a slice at the specified index.
func RemoveAt[T any](slice []T, index int) ([]T, error) {
	if len(slice) == 0 {
		return nil, errors.New("slice is empty")
	}
	if index < 0 || index > len(slice)-1 {
		return nil, errors.New("index out of range")
	}
	firstHalfLen := len(slice[:index])
	secondHalfLen := len(slice[index+1:])
	firstHalf := make([]T, firstHalfLen)
	secondHalf := make([]T, secondHalfLen)
	copy(firstHalf, slice[:index])
	copy(secondHalf, slice[index+1:])
	result := append(firstHalf, secondHalf...)
	return result, nil
}

// Pop removes the last element from a slice.
func Pop[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

// Contains searches a slice for a given value.
func Contains[T comparable](slice []T, searchFor T) bool {
	for _, val := range slice {
		if val == searchFor {
			return true
		}
	}
	return false
}

// Shift removes the first element from a slice.
func Shift[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[1:]
}

// Min returns the lowest entry in the slice.
func Min[T constraints.Ordered](slice []T) (result T) {
	for i, val := range slice {
		if i == 0 || val < result {
			result = val
		}
	}
	return
}

// Max returns the highest entry in the slice.
func Max[T constraints.Ordered](slice []T) (result T) {
	for _, val := range slice {
		if val > result {
			result = val
		}
	}
	return
}
