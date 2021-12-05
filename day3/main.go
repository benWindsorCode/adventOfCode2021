package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

func GetInput() ([]string, error){
	file, err := os.Open("input.txt")

	var result []string

	if err != nil {
		return result, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := scanner.Text()

		result = append(result, x)
	}

	return result, scanner.Err()
}

func Part1(inputs []string) () {
	// Assumption that inputs is even number... its fine for this puzzle input
	threshold := len(inputs)/2
	len := len(inputs[0])
	oneCounts := make([]int, len)

	for _, input := range inputs {
		for i, char := range input {
			value, _ := strconv.Atoi(string(char))
			if value == 1 {
				oneCounts[i] = oneCounts[i] + 1
			}
		}
	}

	var gamma, epsilon string

	// Gamma is most common digit, epsilon is least common digit
	for _, count  := range oneCounts {
		if count > threshold {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	// Convert from binary to decimal
	gammaParsed, errGamma := strconv.ParseInt(gamma, 2, 64)
	epsilonParsed, errEpsilon := strconv.ParseInt(epsilon, 2, 64)

	if errGamma != nil || errEpsilon != nil {
		log.Fatal(errGamma)
		log.Fatal(errEpsilon)
	}

	fmt.Println(gammaParsed * epsilonParsed)
}

func Part2(inputs []string) () {

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
