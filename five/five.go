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

	Ranges []Range
}

type Range struct {
	DstStart int
	SrcStart int
	Length   int
}

func (m *Map) CorrespondingTo(num int) int {
	for _, r := range m.Ranges {
		if r.SrcStart <= num && num <= r.SrcStart+r.Length {
			res := r.DstStart + (num - r.SrcStart)
			return res
		}
	}
	return num
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
		Src:    src,
		Dst:    dst,
		Ranges: []Range{},
	}
	for _, line := range lines[1:] {
		spl := strings.Split(line, " ")
		if len(spl) != 3 {
			return nil, fmt.Errorf("could not parse values from %+v", spl)
		}
		// important but weird, it is destination start, source start, length
		// so {dst, src, len} rather than {src, dst, len}
		dststr, srcstr, lengthstr := spl[0], spl[1], spl[2]
		dststart, err := strconv.Atoi(dststr)
		if err != nil {
			return nil, err
		}
		srcstart, err := strconv.Atoi(srcstr)
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(lengthstr)
		if err != nil {
			return nil, err
		}
		m.Ranges = append(m.Ranges, Range{
			DstStart: dststart,
			SrcStart: srcstart,
			Length:   length,
		})
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
		location := MatchSeedToLocation(seed, maps)
		locations = append(locations, location)
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

func FertilizerPartTwo(input string) (*int, error) {
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
	seeds := *seedsToPlant // deref to index
	for i := 0; i < len(seeds)-1; i += 2 {
		seedstart := seeds[i]
		seedlength := seeds[i+1]
		for j := 0; j < seedlength; j++ {
			location := MatchSeedToLocation(seedstart+j, maps)
			locations = append(locations, location)
		}
	}

	min := Min(locations...)

	return &min, nil
}

func MatchSeedToLocation(seed int, maps Maps) int {
	src := "seed"
	dst := maps[src].Dst
	val := maps[src].CorrespondingTo(seed)
	for {
		if dst == "location" {
			break
		}
		src = dst
		dst = maps[src].Dst
		val = maps[src].CorrespondingTo(val)
	}
	return val
}
