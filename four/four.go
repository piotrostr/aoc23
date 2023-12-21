package four

import (
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	WinningNums []int
	Nums        []int
}

func ArrOfStringToArrOfInt(arr []string) (*[]int, error) {
	new := []int{}
	for _, item := range arr {
		if item == "" {
			continue
		}
		asint, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		new = append(new, asint)
	}
	return &new, nil
}

func Scratchcards(input string) (*int, error) {
	points := 0
	cards := map[string]Card{}

	lines := strings.Split(input, "\n")
	// nums := map[string]
	for _, line := range lines {
		spl := strings.Split(line, ":")
		if len(spl) != 2 {
			return nil, fmt.Errorf("faulty line (no ':'): %s", line)
		}
		index, values := spl[0], spl[1]
		spl = strings.Split(values, "|")
		if len(spl) != 2 {
			return nil, fmt.Errorf("faulty card values (no '|'): %s", line)
		}
		winningNumsStr := strings.Split(spl[0], " ")
		numsStr := strings.Split(spl[1], " ")
		winningNums, err := ArrOfStringToArrOfInt(winningNumsStr)
		if err != nil {
			return nil, err
		}
		nums, err := ArrOfStringToArrOfInt(numsStr)
		if err != nil {
			return nil, err
		}
		cards[index] = Card{
			WinningNums: *winningNums,
			Nums:        *nums,
		}
	}

	for _, card := range cards {
		multiplier := 1
		containedAny := false
		for _, num := range card.Nums {
			if Contains(card.WinningNums, num) {
				if containedAny {
					multiplier = multiplier * 2
				}
				containedAny = true
			}
		}
		if containedAny {
			points += multiplier
		}
	}

	return &points, nil
}

func Contains(arr []int, item int) bool {
	for _, arritem := range arr {
		if arritem == item {
			return true
		}
	}
	return false
}
