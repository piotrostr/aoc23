package one

import (
	"fmt"
	"strconv"
	"strings"
)

func Trebuchet(input string) (*int64, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("Trebuchet(%s): empty input", input)
	}
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

func CalibrateLine(line string) string {
	cornerCases := map[string]string{
		"oneight":   "18",
		"threeight": "38",
		"fiveight":  "58",
		"eightree":  "83",
		"nineight":  "98",
		"sevenine":  "79",
		"twone":     "21",
	}
	for {
		firstMatchIndex := -1
		firstMatchKey := ""
		for k := range cornerCases {
			idx := strings.Index(line, k)
			if idx == -1 {
				continue
			}
			if firstMatchIndex == -1 {
				firstMatchIndex = idx
				firstMatchKey = k
			}
			if idx < firstMatchIndex {
				firstMatchIndex = idx
				firstMatchKey = k
			}
		}
		if firstMatchIndex == -1 {
			break
		}
		// the only thing that we need to ensure is that if we are replacing a
		// number it cannot break `oneight`'s etc
		line = strings.Replace(line, firstMatchKey, cornerCases[firstMatchKey], 1)
	}

	conversionSchema := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for {
		firstMatchIndex := -1
		firstMatchKey := ""
		for k := range conversionSchema {
			idx := strings.Index(line, k)
			if idx == -1 {
				continue
			}
			if firstMatchIndex == -1 {
				firstMatchIndex = idx
				firstMatchKey = k
			}
			if idx < firstMatchIndex {
				firstMatchIndex = idx
				firstMatchKey = k
			}
		}
		if firstMatchIndex == -1 {
			break
		}
		// the only thing that we need to ensure is that if we are replacing a
		// number it cannot break `oneight`'s etc
		line = strings.Replace(line, firstMatchKey, conversionSchema[firstMatchKey], 1)
	}

	return line
}

func Calibrate(input string) string {
	res := []string{}
	lines := strings.Split(input, "\n")
	for l := 0; l < len(lines); l++ {
		line := lines[l]
		calibratedLine := CalibrateLine(line)
		res = append(res, calibratedLine)
	}
	calibratedInput := strings.Join(res, "\n")
	return calibratedInput
}

func CalibratedTrebuchet(input string) (*int64, error) {
	// transform the input string before trebuchet
	calibratedInput := Calibrate(input)
	return Trebuchet(calibratedInput)
}
