package three_test

import (
	"testing"

	"github.com/piotrostr/aoc23/three"
)

func TestGearRatios(t *testing.T) {
	input := `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..`

	res, err := three.GearRatios(input)
	if err != nil {
		t.Fatal(err)
	}

	expected := 4361
	if *res != expected {
		t.Fatalf("GearRatios(%s) => %d != %d", input, *res, expected)
	}
}
