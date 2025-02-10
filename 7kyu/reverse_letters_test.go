package sevenkyu

import (
	"testing"
)

func TestReverseLetters(t *testing.T) {
	result := ReverseLetters("28cba")
	expected := "abc"

	if expected != result {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
