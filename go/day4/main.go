package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func GetInput() ([]int, [100][5][5]int, error){
	// First line of file is values called in bingo, followed by 100 lots of 5x5 bingo grids
	file, err := os.Open("input.txt")

	var lines []string
	var values []int
	var grids [100][5][5]int

	if err != nil {
		return values, grids, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		x := scanner.Text()

		lines = append(lines, x)
	}

	// First line of file is array of positions
	for _, posStr := range strings.Split(lines[0], ",") {
		posInt, _ := strconv.Atoi(posStr)
		values = append(values, posInt)
	}

	// Each grid is five by five with newline in between, starting from 3rd line of file
	gridNum := 0
	for lineNum := 2; lineNum <= len(lines) - 5; lineNum = lineNum + 6 {
		var grid [5][5]int
		for row := 0; row < 5; row++ {
			line := lines[lineNum + row]
			rowPos := 0
			var rowInts [5]int

			for _, gridValStr := range strings.Fields(line) {
				gridVal, _ := strconv.Atoi(gridValStr)
				rowInts[rowPos] = gridVal
				rowPos++
			}

			grid[row] = rowInts
		}

		grids[gridNum] = grid
		gridNum++
	}

	return values, grids, scanner.Err()
}

func evaluateGrid(value int, grid [5][5]int, gridStatus [5][5]int) ([5][5]int){
	// Given a value called, a grid, and current evaluation status of grid, return updated evaluation status
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if grid[row][col] == value {
				gridStatus[row][col] = 1
			}
		}
	}

	return gridStatus
}

func isGridWinning(gridStatus [5][5]int) (bool) {
	// Check all rows
	for row := 0; row < 5; row++ {
		sum := gridStatus[row][0] + gridStatus[row][1] + gridStatus[row][2] + gridStatus[row][3] + gridStatus[row][4]
		if sum == 5 {
			return true
		}
	}

	// Check all cols
	for col := 0; col < 5; col++ {
		sum := gridStatus[0][col] + gridStatus[1][col] + gridStatus[2][col] + gridStatus[3][col] + gridStatus[4][col]
		if sum == 5 {
			return true
		}
	}

	return false
}

func uncalledSum(grid [5][5]int, gridStatus [5][5]int) (int) {
	// Sum the uncalled numbers of the winning grid
	sum := 0

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if gridStatus[row][col] == 0 {
				sum = sum + grid[row][col]
			}
		}
	}

	return sum
}

func Part1(values []int, grids [100][5][5]int) {
	// Record a '1' where each grid val gets called
	var gridStatuses [100][5][5]int

	for _, value := range values {
		for i, grid := range grids {
			currentStatus := gridStatuses[i]
			gridStatuses[i] = evaluateGrid(value, grid, currentStatus)

			if isGridWinning(gridStatuses[i]) {
				uncalledSum := uncalledSum(grids[i], gridStatuses[i])
				answer := uncalledSum * value
				fmt.Println(answer)
				return
			}
		}
	}
}

func Part2(values []int, grids [100][5][5]int) {
	// Record a '1' where each grid val gets called
	var gridStatuses [100][5][5]int

	var winners []int
	var winLast int
	var winLastValue int
	var winLastStatus [5][5]int

	// This is not super pretty nested logic
	for _, value := range values {
		for i, grid := range grids {
			currentStatus := gridStatuses[i]
			gridStatuses[i] = evaluateGrid(value, grid, currentStatus)

			if isGridWinning(gridStatuses[i]) {
				alreadyWon := false
				for _, winner := range winners {
					if winner == i {
						alreadyWon = true
					}
				}

				if !alreadyWon {
					winners = append(winners, i)
					winLast = i
					winLastValue = value
					winLastStatus = gridStatuses[i]
				}
			}
		}
	}

	uncalledSum := uncalledSum(grids[winLast], winLastStatus)
	answer := uncalledSum * winLastValue
	fmt.Println(answer)
}

func main() {
	values, grids, err := GetInput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1")
	Part1(values, grids)

	fmt.Println("Part 2")
	Part2(values, grids)
}
