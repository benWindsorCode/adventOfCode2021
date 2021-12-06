package main

import(
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	var input = [][2][2]int{
		{{0,9}, {5,9}},
		{{8,0}, {0,8}},
		{{9,4}, {3,4}},
		{{2,1}, {2,1}},
		{{7,0}, {7,4}},
		{{6,4}, {2,0}},
		{{0,9}, {2,9}},
		{{3,4}, {1,4}},
		{{0,0}, {8,8}},
		{{5,5}, {8,2}},
	}

	result := Part1(input)

	if result != 5 {
		t.Errorf(fmt.Sprintf("found: %d intersections, expected 5", result))
	}
}
