package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
	"strings"
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
		vals := strings.Split(scanner.Text(), ",")

		for _, val := range vals {
			valInt, _ := strconv.Atoi(val)
			result = append(result, valInt)
		}
	}

	return result, scanner.Err()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Part1(inputs []int) {
	// Snippet from stack overflow - why doesn't go have consts for this...
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	min, max := MaxInt, MinInt

	for _, val := range inputs {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	minFuel := MaxInt
	for i := min; i <= max; i++ {
		totalFuel := 0

		for _, val := range inputs {
			totalFuel = totalFuel + Abs(val - i)
		}

		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}

	fmt.Println(minFuel)
}

func Part2(inputs []int) {
	// Snippet from stack overflow - why doesn't go have consts for this...
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	min, max := MaxInt, MinInt

	for _, val := range inputs {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	minFuel := MaxInt
	for i := min; i <= max; i++ {
		totalFuel := 0
		for _, val := range inputs {
			dist := Abs(val - i)
			fuelCost := dist * (dist+1)/2
			totalFuel = totalFuel + fuelCost
		}

		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}

	fmt.Println(minFuel)
}

func main() {
	inputs, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1")
	Part1(inputs)

	fmt.Println("Part 2")
	Part2(inputs)
}
