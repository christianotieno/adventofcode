package day2

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)


func PartOne(input string) (int, error) {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	total := 0

	for _, r := range ranges {
		rangeParts := strings.Split(r, "-")
		if len(rangeParts) != 2 {
			return 0, fmt.Errorf("invalid range format: %s", r)
		}

		start, err1 := strconv.Atoi(rangeParts[0])
		end, err2 := strconv.Atoi(rangeParts[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid range values: %s", r)
		}
		for n := start; n <= end; n++ {
			if utils.IsInvalidID(n) {
				total += n
			}
		}
	}

	return total, nil
}

func PartTwo(input string) (int, error) {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	total := 0

	for _, r := range ranges {
		rangeParts := strings.Split(r, "-")
		if len(rangeParts) != 2 {
			return 0, fmt.Errorf("invalid range format: %s", r)
		}

		start, err1 := strconv.Atoi(rangeParts[0])
		end, err2 := strconv.Atoi(rangeParts[1])
		if err1 != nil || err2 != nil {
			return 0, fmt.Errorf("invalid range values: %s", r)
		}

		for n := start; n <= end; n++ {
			if utils.IsRepeatedSequence(n) {
				total += n
			}
		}
	}

	return total, nil
}



