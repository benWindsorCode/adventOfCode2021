package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
)

type lanternfish struct {
	new bool
	timer int
}

func NewLanternfish() *lanternfish {
	lanternfish := lanternfish{new: true, timer: 8}

	return &lanternfish
}

func AgeLanternfish(fish *lanternfish) ([]*lanternfish) {
	// If timer is zero, return a new lanternfish too, otherwise just update current

	var newFish []*lanternfish
	if fish.timer == 0 {
		// regardless of current value it must always be false next
		fish.new = false
		fish.timer = 6

		baby := NewLanternfish()
		newFish = append(newFish, baby)
		return newFish
	}

	fish.timer = fish.timer - 1
	return newFish
}

func RunDay(fish []*lanternfish) ([]*lanternfish) {
	var allFish []*lanternfish

	for _, currentFish := range fish {
		potentialNewFish := AgeLanternfish(currentFish)

		if len(potentialNewFish) == 1 {
			allFish = append(allFish, potentialNewFish[0])
		}

		allFish = append(allFish, currentFish)
	}

	return allFish
}

func CreateFish(timers []int) ([]*lanternfish) {
	var fish []*lanternfish

	for _, timer := range timers {
		newFish := lanternfish{new: false, timer: timer}
		fish = append(fish, &newFish)
	}

	return fish
}

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

func Part1(inputs []int) {
	fishList := CreateFish(inputs)

	for i := 0; i < 80; i++ {
		fishList = RunDay(fishList)
	}

	fmt.Println(len(fishList))
}

func RunMapDay(fishMap map[int]int) (map[int]int) {
	updatedFish := make(map[int]int)

	for i := 0; i <= 8; i++ {
		updatedFish[i] = 0
	}

	for key, val := range fishMap {
		// all with timer zero, go back as timer 6 and create a load of new with timer 8
		if key == 0 {
			updatedFish[8] = updatedFish[8] + val
			updatedFish[6] = updatedFish[6] + val
		} else {
			// otherwise they just decrease their timer
			updatedFish[key-1] = updatedFish[key-1] + val
		}
	}

	return updatedFish
}

func Part2(inputs []int) {
	// For performance reasons try instead storing map of { timer := num fish }
	fishMap := make(map[int]int)

	for _, timer := range inputs {
		if val, ok := fishMap[timer]; ok {
			fishMap[timer] = val + 1
		} else {
			fishMap[timer] = 1
		}
	}

	for i := 0; i < 256; i++ {
		fishMap = RunMapDay(fishMap)
	}

	sum := 0
	for _, val := range fishMap {
		sum = sum + val
	}

	println(sum)
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
