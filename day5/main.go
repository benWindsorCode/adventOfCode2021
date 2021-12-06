package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
	"time"
)

func ProcessPair(pairStr string) ([2]int) {
	// Given a string pair "a,b" return two ints: a, b
	xy := strings.Split(pairStr, ",")
	x, y := xy[0], xy[1]

	xInt, _ := strconv.Atoi(x)
	yInt, _ := strconv.Atoi(y)

	pair := [2]int{xInt, yInt}
	return pair
}

func GetInput() ([][2][2]int, error){
	file, err := os.Open("input.txt")

	var result [][2][2]int

	if err != nil {
		return result, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		startEnd := strings.Split(scanner.Text(), " -> ")
		startStr, endStr := startEnd[0], startEnd[1]
		start := ProcessPair(startStr)
		end := ProcessPair(endStr)

	    line := [2][2]int{start, end}

		result = append(result, line)
	}

	return result, scanner.Err()
}

func Max(first int, second int) (int) {
	if first > second {
		return first
	}

	return second
}

func Min(first int, second int) (int) {
	if first < second {
		return first
	}

	return second
}

func GetVerticalLinePoints(x int, yStart int, yEnd int) ([][2]int) {
	yMin := Min(yStart, yEnd)
	yMax := Max(yStart, yEnd)

	var points [][2]int
	for y := yMin; y <= yMax; y++ {
		point := [2]int{x, y}
		points = append(points, point)
	}

	return points
}

func GetHorizontalLinePoints(y int, xStart int, xEnd int) ([][2]int) {
	xMin := Min(xStart, xEnd)
	xMax := Max(xStart, xEnd)

	var points [][2]int
	for x := xMin; x <= xMax; x++ {
		point := [2]int{x, y}
		points = append(points, point)
	}

	return points
}

func GetVerticalOrHorizontal(lines [][2][2]int) ([][2][2]int) {
	var verticalOrHorizontal [][2][2]int

	for _, line := range lines {
		x0, y0, x1, y1 := line[0][0], line[0][1], line[1][0], line[1][1]

		if x0 == x1 || y0 == y1 {
			verticalOrHorizontal = append(verticalOrHorizontal, line)
		}
	}

	return verticalOrHorizontal
}

func Part1(inputs [][2][2]int) (int) {
	// Filter down to all vertical/horizontal lines
	verticalOrHorizontal := GetVerticalOrHorizontal(inputs)

	// Track mapping of point to how many times point occurs
	linePointsCounts := make(map[[2]int]int)
	for _, line := range verticalOrHorizontal {
		x0, y0, x1, y1 := line[0][0], line[0][1], line[1][0], line[1][1]

		var points [][2]int
		if x0 == x1 {
			points = GetVerticalLinePoints(x0, y0, y1)
		} else if y0 == y1 {
			points = GetHorizontalLinePoints(y0, x0, x1)
		}

		for _, point := range points {
			if val, ok := linePointsCounts[point]; ok {
				linePointsCounts[point] = val + 1
			} else {
				linePointsCounts[point] = 1
			}
		}
	}

	// Per problem: count up all times point is touched >= 2 times
	sum := 0
	for _, count := range linePointsCounts {
		if count >= 2 {
			sum++
		}
	}

	return sum
}

func main() {
	start := time.Now()
	inputs, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1")
	res := Part1(inputs)
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Printf("Part 1 took: %s", elapsed)
}
