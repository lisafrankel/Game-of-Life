package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// read input file
func processInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	checkErr(err)
	return string(data)
}

// take input string and put into a 2d bool array 
func formatDat(inital_state string, row int, column int) [][]bool {
	row += 2
	column += 2
	grid := make([][]bool, row)
	for i := range grid {
		grid[i] = make([]bool, column)
	}

	ind := 0
	for i := 1; i < row - 1; i++ {
		for j := 1; j < column - 1; j++ {
			if inital_state[ind] == '\n' {
				ind++
			}
			if inital_state[ind] == '-' {
				grid[i][j] = false
			}
			if inital_state[ind] == '*' {
				grid[i][j] = true
			}
			ind++
		}
	}
	return grid
}

func printGrid(grid [][]bool, row int, column int) {
	for i := 1; i < row + 1; i++ {
		for j := 1; j < column + 1; j++ {
			if grid[i][j] == false {
				fmt.Print(" ")
			}
			if grid[i][j] == true {
				fmt.Print("*")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("===========================\n")
}

type Cell struct {
	xCord int
	yCord int
}

//  returns the number of alive cells surrounding the cell at location (i, j)
func numOfAliveNeighbors(grid [][]bool, i int, j int, row int, column int) int {
	numOfAlive := 0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if grid[i + x][j + y] == true && (!(x == 0 && y ==0))  {
				numOfAlive++
			}
		}
	}
	return numOfAlive
}

// find all locations where a new alive cell should be placed
func findNewAliveCells(grid [][]bool, row int, column int) []Cell {
	newAliveCells := make([]Cell, 0, row * column)
	for i := 1; i < row + 1; i++ {
		for j := 1; j < column + 1; j++ {
			if grid[i][j] == false{
				numOfAlive := numOfAliveNeighbors(grid, i, j, row, column)
				if numOfAlive == 3 {
					//fmt.Print("(", i, ", ", j, "): ", numOfAlive, "\n")
					newCell := Cell{}
					newCell.xCord = i
					newCell.yCord = j
					newAliveCells = append(newAliveCells, newCell)
				}
			}
		}
	}
	return newAliveCells
}

// find all locations where a new dead cell should be placed
func findNewDeadCells(grid [][]bool, row int, column int) []Cell {
	newDeadCells := make([]Cell, 0, row * column)
	for i := 1; i < row + 1; i++ {
		for j := 1; j < column + 1; j++ {
			if grid[i][j] == true {
				numOfAlive := numOfAliveNeighbors(grid, i, j, row, column)
				if numOfAlive > 3 || numOfAlive < 2 {
					//fmt.Print("(", i, ", ", j, "): ", numOfAlive, "\n")
					newCell := Cell{}
					newCell.xCord = i
					newCell.yCord = j
					newDeadCells = append(newDeadCells, newCell)
				}
			}
		}
	}
	return newDeadCells
}

// adds new alive and dead cells to grid
func updateGrid(grid [][]bool, newAliveCells []Cell, newDeadCells []Cell) [][]bool {
	updatedGrid := grid
	for i := 0; i < len(newAliveCells); i++ {
		updatedGrid[newAliveCells[i].xCord][newAliveCells[i].yCord] = true
	}
	for i := 0; i < len(newDeadCells); i++ {
		updatedGrid[newDeadCells[i].xCord][newDeadCells[i].yCord] = false
	}
	return updatedGrid
}

// where next generation is determined
func evolve(grid [][]bool, row int, column int) [][]bool{
	newAliveCells := findNewAliveCells(grid, row, column)
	newDeadCells := findNewDeadCells(grid, row, column)
	updatedGrid := updateGrid(grid, newAliveCells, newDeadCells)
	return updatedGrid
}
	
func main() {
	
	//"/Users/Lisa/Desktop/HOME/go/src/gol/life.txt"
	filename := os.Args[1]
	inital_state := processInput(filename)

	// get # rows and columns
	row := 0
	column := -1
	doneCountingColumns := false
	for i:= 0; i < len(inital_state); i++ {
		if !doneCountingColumns {
			column++
		}
		if inital_state[i] == '\n' {
			row++
			doneCountingColumns = true
		}
	}

	grid := formatDat(inital_state, row, column)
	fmt.Print("Inital world\n")
	printGrid(grid, row, column)
	for i := 0; i < 10; i++ {
		fmt.Print("Generation: ", i+1, "\n")
		updatedGrid := evolve(grid, row, column)
		printGrid(updatedGrid, row, column)

	}

}
