package day3

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)


func PartOne(input string) (int, error) {
    total := 0

    for bank := range strings.FieldsSeq(input) {
        maxJoltage := 0
        for i := 0; i < len(bank); i++ {
            for j := i + 1; j < len(bank); j++ {
                joltage, err := strconv.Atoi(string(bank[i]) + string(bank[j]))
                if err != nil {
                    return 0, fmt.Errorf("invalid digit in bank: %v", err)
                }
                maxJoltage = utils.Max(maxJoltage, joltage)
            }
        }
        total += maxJoltage
    }

    return total, nil
}