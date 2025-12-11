package day2

import "testing"

const testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestPartOne(t *testing.T) {
	expected := 1227775554
	result, err := PartOne(testInput)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
