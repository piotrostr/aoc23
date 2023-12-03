package three

import (
	"fmt"
	"strconv"
	"strings"
)

// GearRatios parses the numbers from the engine schematic, each number
// containing an array of set of coordinates '123': [[i1, j1], [i2, j1], ...] of its
// coordinates
// it then checks if then numbers are valid, writing the results to another
// mapping lastly, the numbers can be processed into a sum if they are valid,
// providing the solution to the AOC problem
func GearRatios(input string) (*int, error) {
	numberLocations := map[string][][]int{}
	lines := strings.Split(input, "\n")
	for i := range lines {
		num := ""
		locations := [][]int{}
		for j := range lines[i] {
			if IsNum(string(lines[i][j])) {
				num += string(lines[i][j])
				locations = append(locations, []int{i, j})
			} else {
				// write res
				if num != "" && len(locations) > 0 {
					numberLocations[num] = locations
				}

				// reset num
				num = ""
				locations = [][]int{}
			}
		}
	}
	fmt.Printf("%+v", numberLocations)
	return nil, nil
}

func IsNum(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}
	return false
}
