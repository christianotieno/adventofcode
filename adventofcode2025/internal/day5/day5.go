package day5

import (
	"adventofcode2025/utils"
	"sort"
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges, ids, err := utils.ParseInput(lines)
	if err != nil {
		return 0, err
	}

	count := 0
	for _, id := range ids {
		if utils.IsFresh(id, ranges) {
			count++
		}
	}
	return count, nil
}

func PartTwo(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ranges, _, err := utils.ParseInput(lines)
	if err != nil {
		return 0, err
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := make([]utils.Range, 0, len(ranges))
	for _, r := range ranges {
		if len(merged) == 0 || merged[len(merged)-1].End < r.Start-1 {
			merged = append(merged, r)
		} else if r.End > merged[len(merged)-1].End {
			merged[len(merged)-1].End = r.End
		}
	}

	total := 0
	for _, r := range merged {
		total += r.End - r.Start + 1
	}

	return total, nil
}

