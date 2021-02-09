package slices

import (
	"errors"
	"testing"
)

func TestSomeTrue(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := true
	actual, ok := Some(input, func(index int, value interface{}) bool {
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

func TestSomeFalse(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := false
	actual, ok := Some(input, func(index int, value interface{}) bool {
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

func TestSomeNonSlice(t *testing.T) {
	input := "testing"
	expect := errors.New("Some requires a slice")
	_, ok := Some(input, func(index int, value interface{}) bool {
		if val, isString := value.(string); isString {
			return len(val) == 0
		}
		return false
	})

	if ok.Error() != expect.Error() {
		failTest(t, expect, ok)
	}
}
