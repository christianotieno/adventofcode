package day3

import "testing"

const testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestPartOne(t *testing.T) {
	expected := 357

	result, err := PartOne(testInput)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

