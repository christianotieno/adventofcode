package utils

import (
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func IsInvalidID(n int) bool {
	if n < 10 {
		return false
	}
	
	s := strconv.Itoa(n)
	length := len(s)
	
	if length%2 != 0 {
		return false
	}
	
	half := length / 2
	return s[:half] == s[half:]
}

func IsRepeatedSequence(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue 
		}

		pattern := s[:patternLen]
		isValid := true

		for i := patternLen; i < length; i += patternLen {
			if s[i:i+patternLen] != pattern {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}

func Max(a, b int) int {
    if a > b { return a }
    return b
}

func IsFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}

func ParseInput(lines []string) ([]Range, []int, error) {
	var ranges []Range
	var ids []int
	processingRanges := true

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			processingRanges = false
			continue
		}

		if processingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				continue
			}
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				ranges = append(ranges, Range{Start: start, End: end})
			}
		} else {
			if id, err := strconv.Atoi(line); err == nil {
				ids = append(ids, id)
			}
		}
	}

	return ranges, ids, nil
}

func GetFreshIDs(ranges []Range) map[int]bool {
	freshIDs := make(map[int]bool)

	for _, r := range ranges {
		for i := r.Start; i <= r.End; i++ {
			freshIDs[i] = true
		}
	}

	return freshIDs
}
