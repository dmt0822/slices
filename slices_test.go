package slices

import (
	"errors"
	"fmt"
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

func TestPop(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"this", "is", "a", "test"}
	actual, ok := Pop(input)

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
			return
		}
	}
}

func TestPopNonSlice(t *testing.T) {
	input := "this is not a slice"
	expect := "Pop requires a slice"
	actual, ok := Pop(input)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, actual)
		return
	}
}

func TestRemoveAt(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"this", "a", "test", "slice"}
	actual, ok := RemoveAt(input, 1)

	if ok != nil {
		failTest(t, expect, ok)
		return
	}

	for index, actualVal := range actual {
		if actualVal != expect[index] {
			failTest(t, expect, actual)
			return
		}
	}
}

func TestRemoveAtNonSlice(t *testing.T) {
	input := "i'm not a slice"
	expect := "RemoveAt requires a slice"
	_, ok := RemoveAt(input, 1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtEmptySlice(t *testing.T) {
	input := []string{}
	expect := "Slice is empty"
	_, ok := RemoveAt(input, 1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtIndexTooLow(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expect := "Index out of range"
	_, ok := RemoveAt(input, -1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtIndexTooHigh(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expect := "Index out of range"
	_, ok := RemoveAt(input, 4)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestShift(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"is", "a", "test", "slice"}
	actual, ok := Shift(input)

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
			return
		}
	}
}

func TestShiftNonSlice(t *testing.T) {
	input := "this is not a slice"
	expect := "Shift requires a slice"
	actual, ok := Shift(input)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, actual)
		return
	}
}

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

func failTest(t *testing.T, expect interface{}, actual interface{}) {
	t.Errorf(fmt.Sprintf("\nExpected: %v\nActual: %v", expect, actual))
}
