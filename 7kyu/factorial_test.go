package sevenkyu

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	result := Factorial(5)
	expected := 120

	if expected != result {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
