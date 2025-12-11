package day5

import "testing"

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

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
    expected := 14 // 3-5,10-14,16-20,12-18 covers 3,4,5,10,11,12,13,14,15,16,17,18,19,20
	result, err := PartTwo(testInput)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}