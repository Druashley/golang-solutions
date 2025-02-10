package sevenkyu

import (
	"testing"
)

func TestLastChcker(t *testing.T) {
	result := LastChecker("abc", "bc")
	expected := true

	if expected != result {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
