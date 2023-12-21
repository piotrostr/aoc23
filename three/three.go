package three

import (
	"strconv"
	"strings"
)

func GearRatios(input string) (*int, error) {
	numberLocations := map[string][][]int{}
	lines := strings.Split(input, "\n")

	// Parse the schematic for numbers and their locations
	for i, line := range lines {
		num := ""
		locations := [][]int{}
		for j, char := range line {
			if char >= '0' && char <= '9' { // Check if character is a digit
				num += string(char)
				locations = append(locations, []int{i, j})
			} else {
				if num != "" {
					numberLocations[num] = append(numberLocations[num], locations...)
				}
				num = ""
				locations = [][]int{}
			}
		}
		if num != "" {
			numberLocations[num] = append(numberLocations[num], locations...)
		}
	}

	sumOfPartNums := 0
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	// Check if numbers are part numbers
	for numstr, arrOfCords := range numberLocations {
		isPartNum := false

		for _, coords := range arrOfCords {
			x, y := coords[1], coords[0]

			// Check all adjacent positions for symbols
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if dx == 0 && dy == 0 {
						continue // Skip the number itself
					}
					nx, ny := x+dx, y+dy
					if nx >= 0 && ny >= 0 && nx <= maxX && ny <= maxY {
						if IsSymbol(lines[ny][nx]) {
							isPartNum = true
							break
						}
					}
				}
				if isPartNum {
					break
				}
			}
		}

		if isPartNum {
			num, err := strconv.Atoi(numstr)
			if err != nil {
				return nil, err
			}
			sumOfPartNums += num
		}
	}

	return &sumOfPartNums, nil
}

func IsSymbol(char byte) bool {
	return char != '.' && (char < '0' || char > '9')
}
