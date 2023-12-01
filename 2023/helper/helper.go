package helper;

import (
	"reflect"
	"testing"
)

func AssertEq[T any](t testing.TB, left, right T) {
	t.Helper()

	if !reflect.DeepEqual(left, right) {
		t.Errorf("\nLeft: '%v'\nRight: '%v'", left, right)
	}
}

func Assert[T bool](t testing.TB, value T) {
	t.Helper()

	if !value {
		t.Errorf("Failed")
	}
}
