package eightkyu

import (
	"testing"
)

func TestHoopCount(t *testing.T) {
	result1 := HoopCount(9)
	expected := "Keep at it until you get it"

	if result1 != expected {
		t.Errorf("Expected %s, but got %s", expected, result1)
	}

	result2 := HoopCount(10)
	expected2 := "Great, now move on to tricks"
	if result2 != expected2 {
		t.Errorf("Expected %s, but got %s", expected2, result2)
	}
}
