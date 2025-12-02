package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// day1.go: part one https://adventofcode.com/2025/day/1

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	position := 50
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if len(line) < 2 {
			continue
		}

		direction := line[0]
		distanceStr := line[1:]

		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return 0, fmt.Errorf("invalid distance in rotation %q: %w", line, err)
		}

		switch direction {
		case 'R':
			position = (position + distance) % 100
		case 'L':
			position = (position - distance + 100) % 100
		default:
			return 0, fmt.Errorf("invalid direction in rotation %q: expected L or R, got %c", line, direction)
		}

		if position == 0 {
			count++
		}
	}

	return count, nil
}
