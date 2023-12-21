package three

import (
	"fmt"
	"strconv"
	"strings"
)

func PushNumberLocation(numberLocations map[string][][]int, numstr string, i int, j int, locations [][]int) {
	index := fmt.Sprintf("%s-%d-%d", numstr, i, j)
	numberLocations[index] = append(numberLocations[index], locations...)
}

func GearRatios(input string) (*int, error) {
	numberLocations := map[string][][]int{}
	lines := strings.Split(input, "\n")

	// Parse the schematic for numbers and their locations
	for i, line := range lines {
		num := ""
		locations := [][]int{}
		for j, char := range line {
			if IsDigit(char) {
				num += string(char)
				locations = append(locations, []int{i, j})
			} else {
				if num != "" {
					PushNumberLocation(numberLocations, num, i, j, locations)
				}
				num = ""
				locations = [][]int{}
			}
		}
		if num != "" {
			PushNumberLocation(numberLocations, num, i, len(line), locations)
		}
	}

	sumOfPartNums := 0
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	asteriskNums := map[string][]string{}

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
							if lines[ny][nx] == '*' {
								index := fmt.Sprintf("%d-%d", nx, ny)
								asteriskNums[index] = append(asteriskNums[index], numstr)
							}
							isPartNum = true
						}
					}
				}
				if isPartNum {
					break
				}
			}
		}

		if isPartNum {
			num, err := strconv.Atoi(strings.Split(numstr, "-")[0])
			if err != nil {
				return nil, err
			}
			sumOfPartNums += num
		}
	}

	return &sumOfPartNums, nil
}

func GearRatiosPartTwo(input string) (*int, error) {
	numberLocations := map[string][][]int{}
	lines := strings.Split(input, "\n")

	// Parse the schematic for numbers and their locations
	for i, line := range lines {
		num := ""
		locations := [][]int{}
		for j, char := range line {
			if IsDigit(char) {
				num += string(char)
				locations = append(locations, []int{i, j})
			} else {
				if num != "" {
					PushNumberLocation(numberLocations, num, i, j, locations)
				}
				num = ""
				locations = [][]int{}
			}
		}
		if num != "" {
			PushNumberLocation(numberLocations, num, i, len(line), locations)
		}
	}

	sumOfPartNums := 0
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	asteriskNums := map[string][]string{}

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
							if lines[ny][nx] == '*' {
								index := fmt.Sprintf("%d-%d", nx, ny)
								// prevent duplicates
								if !strings.Contains(strings.Join(asteriskNums[index], ","), numstr) {
									asteriskNums[index] = append(asteriskNums[index], numstr)
								}
							}
							isPartNum = true
						}
					}
				}
				if isPartNum {
					break
				}
			}
		}

		if isPartNum {
			num, err := strconv.Atoi(strings.Split(numstr, "-")[0])
			if err != nil {
				return nil, err
			}
			sumOfPartNums += num
		}
	}

	sumOfGearRatios := 0
	for _, nums := range asteriskNums {
		if len(nums) == 2 {
			num1, err := strconv.Atoi(strings.Split(nums[0], "-")[0])
			if err != nil {
				return nil, err
			}
			num2, err := strconv.Atoi(strings.Split(nums[1], "-")[0])
			if err != nil {
				return nil, err
			}
			sumOfGearRatios += num1 * num2
		}
	}

	return &sumOfGearRatios, nil
}

func IsSymbol(char byte) bool {
	return char != '.' && (char < '0' || char > '9')
}

func IsDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
