package utils

import "strconv"


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

