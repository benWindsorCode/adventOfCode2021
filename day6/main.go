package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
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
func main() {
	inputs, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(inputs)
}
