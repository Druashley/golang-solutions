package sevenkyu

import (
	"testing"
)

func TestBusNumber(t *testing.T) {
	result := BusNumber([][2]int{{10, 0}, {3, 5}, {5, 8}})
	expected := 5

	if expected != result {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
