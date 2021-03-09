package slices

import (
	"testing"
)

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
