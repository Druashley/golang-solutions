package eightkyu

import "testing"

func TestMove(t *testing.T) {
	result := Move(3, 6)
	expected := 15

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
