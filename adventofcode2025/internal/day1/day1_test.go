package day1

import "testing"

const testInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPartOne(t *testing.T) {
	expected := 3
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 6
	result, err := PartTwo(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
