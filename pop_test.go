package slices

import (
	"testing"
)

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
