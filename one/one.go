package one

import (
	"strconv"
	"strings"
)

func Trebuchet(input string) (*int64, error) {
	var sum int64 = 0
	lines := strings.Split(input, "\n")
	for l := 0; l < len(lines); l++ {
		firstDigitIdx := -1
		lastDigitIdx := -1
		line := lines[l]
		for i := 0; i < len(line); i++ {
			if _, err := strconv.ParseInt(string(line[i]), 10, 0); err == nil {
				if firstDigitIdx == -1 {
					firstDigitIdx = i
				}
				if lastDigitIdx < i {
					lastDigitIdx = i
				}
			}
		}
		elems := []string{
			string(line[firstDigitIdx]),
			string(line[lastDigitIdx]),
		}
		joined := strings.Join(elems, "")
		joinedInt, err := strconv.ParseInt(joined, 10, 0)
		if err != nil {
			return nil, err
		}
		sum += joinedInt
	}

	return &sum, nil
}
