package day4

import "testing"

const testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestPartOne(t *testing.T) {
    expected := 13
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}