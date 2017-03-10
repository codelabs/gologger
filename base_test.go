package logger

import (
	"testing"
)

func assertEquality(t *testing.T, expected interface{}, received interface{}) {
	if expected != received {
		t.Errorf("Expected: %v Received: %v\n", expected, received)
	}
}
