package day4

import (
	"strings"
)

func PartOne(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0, nil
	}

	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]rune, rows)
	for i := range grid {
		grid[i] = []rune(lines[i])
	}

	count := 0

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},          {0, 1},
		{1, -1},  {1, 0}, {1, 1},
	}

	for i := range rows {
		for j := range cols {
			if grid[i][j] != '@' {
				continue
			}

			adjacentRolls := 0
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == '@' {
					adjacentRolls++
				}
			}

			if adjacentRolls < 4 {
				count++
			}
		}
	}

	return count, nil
}