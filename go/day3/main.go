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

func OneCounts(inputs []string) ([]int) {
	// Return number of counts of 1s totalled at each position of input
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

	return oneCounts
}

func Part1(inputs []string) () {
	// Assumption that inputs is even number... its fine for this puzzle input
	threshold := len(inputs)/2
	oneCounts := OneCounts(inputs)

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

func ComputeValue(inputs []string, keepMostCommon bool, defaultVal rune) (int64) {
	// Assumption that inputs is even number... its fine for this puzzle input
	threshold := len(inputs)/2

	currentPos := 0
	remaining := len(inputs)

	for remaining > 1 {
		counts := OneCounts(inputs)

		currentOnes := counts[currentPos]

		var keepVal rune
		if remaining % 2 == 0 && currentOnes == threshold {
			// If tie then we use default as tiebreaker
			keepVal = defaultVal
		} else if keepMostCommon && currentOnes > threshold {
			// If we keep the most common value, and one is most common then keep
			keepVal = '1'
		} else if keepMostCommon && currentOnes <= threshold {
			// If we keep most common and zero is most common then keep zero
			keepVal = '0'
		} else if currentOnes > threshold {
			// If we keep least common and one is most common then keep zero
			keepVal = '0'
		} else {
			// If we keep least common and zero is most common then keep one
			keepVal = '1'
		}

		// Keep only inputs where currentPos char is equal to keepVal
		var inputsNew []string
		for _, input := range inputs {
			currentChar := rune(input[currentPos])
			if currentChar == keepVal {
				inputsNew = append(inputsNew, input)
			}
		}

		inputs = inputsNew
		remaining = len(inputs)
		threshold = remaining/2
		currentPos++
	}

	valParsed, err := strconv.ParseInt(inputs[0], 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	return valParsed
}

func Part2(inputs []string) () {
	keepMostCommon := true
	defaultVal := '1'
	oxygenGeneratorRating := ComputeValue(inputs, keepMostCommon, defaultVal)
	keepMostCommon = false
	defaultVal = '0'
	co2ScrubberRating := ComputeValue(inputs, keepMostCommon, defaultVal)

	fmt.Println(oxygenGeneratorRating * co2ScrubberRating)
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
