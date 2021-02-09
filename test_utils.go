package slices

import (
	"fmt"
	"testing"
)

func failTest(t *testing.T, expect interface{}, actual interface{}) {
	t.Errorf(fmt.Sprintf("\nExpected: %v\nActual: %v", expect, actual))
}
