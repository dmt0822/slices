package slices

import (
	"testing"
)

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
