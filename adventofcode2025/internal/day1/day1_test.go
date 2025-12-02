package day1

import "testing"

func TestPartOne(t *testing.T) {
	expected := 0
	result, err := PartOne("")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
