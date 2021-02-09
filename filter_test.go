package slices

import (
	"errors"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"this", "test", "slice"}
	actual, ok := Filter(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) > 2
		}
		return false
	})

	if ok != nil {
		failTest(t, expect, ok)
		return
	}

	if len(actual) != len(expect) {
		failTest(t, expect, actual)
		return
	}

	for index, actualVal := range actual {
		if actualVal != expect[index] {
			failTest(t, expect, actual)
		}
	}
}

func TestFilterNonSlice(t *testing.T) {
	input := "testing"
	expect := errors.New("Filter requires a slice")
	_, ok := Filter(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) == 0
		}
		return false
	})

	if ok == nil {
		failTest(t, expect, ok)
		return
	}

	if ok.Error() != expect.Error() {
		failTest(t, expect, ok)
	}
}
