package one_test

import (
	"testing"

	"github.com/piotrostr/aoc23/one"
)

func TestOne(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	res, err := one.Trebuchet(input)
	if err != nil {
		t.Fatal(err)
	}
	var expect int64 = 142
	if *res != expect {
		t.Errorf("Trebuchet(%s) = %d; want %d", input, res, expect)
	}
}
