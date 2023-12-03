package two

import (
	"fmt"
	"strconv"
	"strings"
)

var BagContents = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// ProcessLine returns whether the line is valid, its sum and error in case
// something is off
func ProcessLine(line string) (*int, bool, error) {
	gameVsRecordsSplit := strings.Split(line, ":")
	if len(gameVsRecordsSplit) != 2 {
		return nil, false, fmt.Errorf("invalid input %s", line)
	}

	id, err := strconv.Atoi(strings.Split(gameVsRecordsSplit[0], " ")[1])
	if err != nil {
		return nil, false, err
	}
	records := gameVsRecordsSplit[1]

	setsSplit := strings.Split(records, ";")

	for i := range setsSplit {
		set := setsSplit[i]
		if strings.Contains(set, ",") {
			cubes := strings.Split(set, ",")
			for j := range cubes {
				num, color, err := ProcessCube(cubes[j])
				if err != nil {
					return &id, false, err
				}
				if *num > BagContents[*color] {
					return &id, false, nil
				}
			}
		} else {
			num, color, err := ProcessCube(set)
			if err != nil {
				return &id, false, nil
			}
			if *num > BagContents[*color] {
				return &id, false, nil
			}
		}
	}

	return &id, true, nil
}

func ProcessCube(cube string) (*int, *string, error) {
	numVsColor := strings.Split(strings.Trim(cube, " "), " ")
	if len(numVsColor) != 2 {
		fmt.Printf("%+v", numVsColor)
		return nil, nil, fmt.Errorf("could not ProcessCube(%s), len(split) != 2", cube)
	}
	numstr, color := numVsColor[0], numVsColor[1]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		return nil, nil, fmt.Errorf("could not ParseInt from %s", numstr)
	}
	return &num, &color, nil
}

func CubeConundrum(input string) (*int, error) {
	sumOfValidIDs := 0
	lines := strings.Split(input, "\n")
	for i := range lines {
		line := lines[i]
		id, valid, err := ProcessLine(line)
		if err != nil {
			return nil, err
		}
		if valid {
			sumOfValidIDs += *id
		}
	}

	return &sumOfValidIDs, nil
}
