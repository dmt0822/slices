package slices

import (
	"errors"
	"reflect"
)

// FilterInt filters a slice of integers.
func FilterInt(slice []int, callback func(index int, value int) bool) []int {
	var result []int
	for index, value := range slice {
		if callback(index, value) {
			result = append(result, value)
		}
	}
	return result
}

// FilterString filters a slice of strings.
func FilterString(slice []string, callback func(index int, value string) bool) []string {
	var result []string
	for index, value := range slice {
		if callback(index, value) {
			result = append(result, value)
		}
	}
	return result
}

// Filter returns a slice of values where callback evaluates to true.
func Filter(arg interface{}, callback func(index int, value interface{}) bool) (interface{}, error) {
	if isSlice(arg) {
		slice := makeSlice(arg)
		var result []interface{}

		for index, value := range slice {
			if callback(index, value) {
				result = append(result, value)
			}
		}

		return result, nil
	}

	return nil, errors.New("Filter requires a slice")
}

// Every returns true if callback evaluates to true for every value in the slice.
func Every(arg interface{}, callback func(index int, value interface{}) bool) (bool, error) {
	if isSlice(arg) {
		slice := makeSlice(arg)

		for index, value := range slice {
			if !callback(index, value) {
				return false, nil
			}
		}
		return true, nil
	}
	return false, errors.New("Every requires a slice")
}

// Some returns true if callback evaluates to true for any value in the slice.
func Some(arg interface{}, callback func(index int, value interface{}) bool) (bool, error) {
	if isSlice(arg) {
		slice := makeSlice(arg)

		for index, value := range slice {
			if callback(index, value) {
				return true, nil
			}
		}
		return false, nil
	}
	return false, errors.New("Some requires a slice")
}

func isSlice(arg interface{}) bool {
	argValue := reflect.ValueOf(arg)
	return argValue.Kind() == reflect.Slice
}

func makeSlice(arg interface{}) []interface{} {
	argValue := reflect.ValueOf(arg)
	slice := make([]interface{}, argValue.Len())

	for i := 0; i < argValue.Len(); i++ {
		slice[i] = argValue.Index(i).Interface()
	}

	return slice
}
