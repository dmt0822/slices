package slices

import (
	"errors"
	"testing"
)

func TestEveryTrue(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := true
	actual, ok := Every(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) > 0
		}
		return false
	})

	if ok != nil {
		failTest(t, expect, ok)
		return
	}

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestEveryFalse(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := false
	actual, ok := Every(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) == 0
		}
		return false
	})

	if ok != nil {
		failTest(t, expect, ok)
		return
	}

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestEveryNonSlice(t *testing.T) {
	input := "testing"
	expect := errors.New("Every requires a slice")
	_, ok := Every(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) == 0
		}
		return false
	})

	if ok.Error() != expect.Error() {
		failTest(t, expect, ok)
	}
}
