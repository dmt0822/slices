package slices

import (
	"testing"
)

func TestEveryTrue(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := true
	actual := Every(input, func(index int, value string) bool {
		return len(value) > 0
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestEveryFalse(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := false
	actual := Every(input, func(index int, value string) bool {
		return len(value) > 3
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestEveryEmptySlice(t *testing.T) {
	var input []string
	expect := false
	actual := Every(input, func(index int, value string) bool {
		return len(value) > 0
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestFilterWithString(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"this", "test", "slice"}
	actual := Filter(input, func(index int, value string) bool {
		return len(value) > 2
	})

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

func TestFilterWithInt(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6}
	expect := []int{0, 2, 4, 6}
	actual := Filter(input, func(index int, value int) bool {
		return value%2 == 0
	})

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

func TestFilterEmptySlice(t *testing.T) {
	input := []int{}
	expect := []int{}
	actual := Filter(input, func(index, value int) bool {
		return value%2 == 0
	})
	if len(actual) > 0 {
		failTest(t, expect, actual)
		return
	}
}

func TestFilterNilSlice(t *testing.T) {
	var input []int
	actual := Filter(input, func(index int, value int) bool {
		return value == 0
	})
	if len(actual) > 0 {
		failTest(t, []int{}, actual)
		return
	}
}

func TestPop(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"this", "is", "a", "test"}
	actual := Pop(input)

	for index, actualVal := range actual {
		if actualVal != expect[index] {
			failTest(t, expect, actual)
			return
		}
	}
}

func TestPopNilSlice(t *testing.T) {
	var input []string
	actual := Pop(input)

	if actual != nil {
		failTest(t, nil, actual)
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

func TestRemoveAtNilSlice(t *testing.T) {
	var input []string
	expect := "slice is empty"
	_, ok := RemoveAt(input, 1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtEmptySlice(t *testing.T) {
	input := []string{}
	expect := "slice is empty"
	_, ok := RemoveAt(input, 1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtIndexTooLow(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expect := "index out of range"
	_, ok := RemoveAt(input, -1)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestRemoveAtIndexTooHigh(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expect := "index out of range"
	_, ok := RemoveAt(input, 4)

	if ok == nil || ok.Error() != expect {
		failTest(t, expect, ok)
	}
}

func TestShift(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := []string{"is", "a", "test", "slice"}
	actual := Shift(input)

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

func TestShiftNilSlice(t *testing.T) {
	var input []string
	actual := Shift(input)

	if actual != nil {
		failTest(t, nil, actual)
		return
	}
}

func TestSomeTrue(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := true
	actual := Some(input, func(index int, value string) bool {
		return len(value) > 0
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestSomeFalse(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := false
	actual := Some(input, func(index int, value string) bool {
		return len(value) == 0
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestSomeNilSlice(t *testing.T) {
	var input []int
	expect := false
	actual := Some(input, func(index int, value int) bool {
		return value > 0
	})

	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestContainsTrue(t *testing.T) {
	input := []string{"this", "is", "a", "test", "slice"}
	expect := true
	actual := Contains(input, "is")
	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestContainsFalse(t *testing.T) {
	input := []string{"another", "test", "slice"}
	expect := false
	actual := Contains(input, "not found")
	if actual != expect {
		failTest(t, expect, actual)
	}
}

func TestContainsNilSlice(t *testing.T) {
	var input []string
	expect := false
	actual := Contains(input, "not found")
	if actual != expect {
		failTest(t, expect, actual)
	}
}

func failTest(t *testing.T, expect interface{}, actual interface{}) {
	t.Errorf("Expected: %v\nActual: %v\n", expect, actual)
}
