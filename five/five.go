package five

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/piotrostr/aoc23/four"
)

type Maps map[string]*Map

type Map struct {
	Src string
	Dst string

	// Items is a mapping of Src-to-Dst
	// e.g.
	// seed 50 => soil 98
	// seed 51 => soil 99
	Items map[int]int
}

func (m *Map) CorrespondingTo(number int) int {
	if val, ok := m.Items[number]; ok {
		return val
	}
	return number
}

func ParseMap(block string) (*Map, error) {
	lines := strings.Split(block, "\n")
	indexline := lines[0]
	spl := strings.Split(strings.Split(indexline, " ")[0], "-")
	if len(spl) != 3 {
		return nil, fmt.Errorf("could not parse values from %+v", spl)
	}
	src, dst := spl[0], spl[2]
	m := &Map{
		Src:   src,
		Dst:   dst,
		Items: map[int]int{},
	}
	for _, line := range lines[1:] {
		spl := strings.Split(line, " ")
		if len(spl) != 3 {
			return nil, fmt.Errorf("could not parse values from %+v", spl)
		}
		// important but weird, it is destination start, source start, length
		// so {dst, src, len} rather than {src, dst, len}
		dststr, srcstr, lengthstr := spl[0], spl[1], spl[2]
		dst, err := strconv.Atoi(dststr)
		if err != nil {
			return nil, err
		}
		src, err := strconv.Atoi(srcstr)
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(lengthstr)
		if err != nil {
			return nil, err
		}
		for i := 0; i < length; i++ {
			m.Items[src+i] = dst + i
		}
	}
	return m, nil
}

func ParseSeedsToPlant(block string) (*[]int, error) {
	vals := strings.Split(block, ":")[1]
	seedstrs := strings.Split(vals, " ")
	return four.ArrOfStringToArrOfInt(seedstrs)
}

func Fertilizer(input string) (*int, error) {
	maps := Maps{}
	blocks := strings.Split(input, "\n\n")
	seedsblock := blocks[0]
	blocks = blocks[1:]
	seedsToPlant, err := ParseSeedsToPlant(seedsblock)
	if err != nil {
		return nil, err
	}
	for _, block := range blocks {
		m, err := ParseMap(block)
		maps[m.Src] = m
		if err != nil {
			return nil, err
		}
	}

	locations := []int{}
	for _, seed := range *seedsToPlant {
		src := "seed"
		dst := maps[src].Dst
		val := maps[src].CorrespondingTo(seed)
		for {
			if dst == "location" {
				locations = append(locations, maps[src].CorrespondingTo(val))
				break
			}
			src = dst
			dst = maps[src].Dst
			val = maps[src].CorrespondingTo(val)
		}
	}

	min := Min(locations...)

	return &min, nil
}

func Min(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
