package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func GetInput() ([][]string, error){
	file, err := os.Open("input.txt")

	var result [][]string

	if err != nil {
		return result, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := strings.Fields(scanner.Text())
		result = append(result, x)
	}

	return result, scanner.Err()
}

func Part1(inputs [][]string) {
	horiz, depth := 0, 0

	for _, input := range inputs {
		instruction, value := input[0], input[1]
		valueParsed, error := strconv.Atoi(value)

		if error != nil {
			log.Fatal(error)
		}

		switch instruction {
		case "up":
			depth = depth - valueParsed
		case "down":
			depth = depth + valueParsed
		case "forward":
			horiz = horiz + valueParsed
		}
	}

	out := horiz * depth

	fmt.Println(out)
}

func Part2(inputs [][]string) {
	horiz, depth, aim := 0, 0, 0

	for _, input := range inputs {
		instruction, value := input[0], input[1]
		valueParsed, error := strconv.Atoi(value)

		if error != nil {
			log.Fatal(error)
		}

		switch instruction {
		case "up":
			aim = aim - valueParsed
		case "down":
			aim = aim + valueParsed
		case "forward":
			horiz = horiz + valueParsed
			depth = depth + aim * valueParsed
		}
	}

	out := horiz * depth

	fmt.Println(out)
}

func main() {
	inputs, error := GetInput()

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Part 1")
	Part1(inputs)

	fmt.Println("Part 2")
	Part2(inputs)
}
