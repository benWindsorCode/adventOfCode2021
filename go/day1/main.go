package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func GetInput() ([]int, error){
	file, err := os.Open("input.txt")

	var result []int

	if err != nil {
		return result, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return result, err
		}

		result = append(result, x)
	}

	return result, scanner.Err()
}

func Part1(input []int) {

	bigger := 0
	for n := 1; n < len(input); n++ {
		if input[n] > input[n-1] {
			bigger++
		}
	}

	fmt.Println(bigger)
}

func Part2(input []int) {
	// To compare rolling 3 window values: (b + c + d) - (a + b + c) > 0
	// <=> d - a > 0
	// so only need to compare elements with a three value offset not the whole sums
	bigger := 0
	for n := 0; n < len(input) - 3; n++ {
		if input[n+3] - input[n] > 0 {
			bigger++
		}
	}

	fmt.Println(bigger)
}

func main() {
	input, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part1")
	Part1(input)

	fmt.Println("Part2")
	Part2(input)
}
