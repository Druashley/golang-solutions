package eightkyu

import "testing"

func TestSumTwoNumbers(t *testing.T) {
	result := SumTwoNumbers(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
