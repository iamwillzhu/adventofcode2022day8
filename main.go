package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type TreeGridWrapper struct {
	TreeGrid     [][]*Tree
	NumberOfRows int
	NumberOfCols int
}

type Tree struct {
	Height    int
	IsVisible bool
}

func (t *TreeGridWrapper) GetNumberVisibleTrees() int {
	numberVisibleTrees := 0
	currentCol := 0
	currentRow := 0

	// set trees on the edge of grid to be visible
	for currentRow < t.NumberOfRows {
		t.TreeGrid[currentRow][0].IsVisible = true
		t.TreeGrid[currentRow][t.NumberOfCols-1].IsVisible = true
		currentRow += 1
	}

	for currentCol < t.NumberOfCols {
		t.TreeGrid[0][currentCol].IsVisible = true
		t.TreeGrid[t.NumberOfCols-1][currentCol].IsVisible = true
		currentCol += 1
	}

	// set tree to be visible from the left side view
	currentRow = 1
	for currentRow < t.NumberOfRows-1 {
		tallestTreeHeightSoFar := t.TreeGrid[currentRow][0].Height

		currentCol = 1
		for currentCol < t.NumberOfCols-1 {
			currentTree := t.TreeGrid[currentRow][currentCol]

			if currentTree.Height > tallestTreeHeightSoFar {
				tallestTreeHeightSoFar = currentTree.Height
				currentTree.IsVisible = true
			}
			currentCol += 1
		}
		currentRow += 1
	}

	// set tree to be visible from top side view
	currentCol = 1
	for currentCol < t.NumberOfCols-1 {
		tallestTreeHeightSoFar := t.TreeGrid[0][currentCol].Height

		currentRow = 1
		for currentRow < t.NumberOfRows-1 {
			currentTree := t.TreeGrid[currentRow][currentCol]

			if currentTree.Height > tallestTreeHeightSoFar {
				tallestTreeHeightSoFar = currentTree.Height
				currentTree.IsVisible = true
			}

			currentRow += 1
		}
		currentCol += 1
	}

	// set tree to be visible from right side view
	currentRow = 1
	for currentRow < t.NumberOfRows-1 {
		tallestTreeHeightSoFar := t.TreeGrid[currentRow][t.NumberOfCols-1].Height

		currentCol = t.NumberOfCols - 2

		for currentCol > 0 {
			currentTree := t.TreeGrid[currentRow][currentCol]

			if currentTree.Height > tallestTreeHeightSoFar {
				tallestTreeHeightSoFar = currentTree.Height
				currentTree.IsVisible = true
			}
			currentCol -= 1
		}
		currentRow += 1
	}

	// set tree to be visible from bottom side view
	currentCol = 1
	for currentCol < t.NumberOfCols-1 {
		tallestTreeHeightSoFar := t.TreeGrid[t.NumberOfRows-1][currentCol].Height

		currentRow = t.NumberOfRows - 2
		for currentRow > 0 {
			currentTree := t.TreeGrid[currentRow][currentCol]

			if currentTree.Height > tallestTreeHeightSoFar {
				tallestTreeHeightSoFar = currentTree.Height
				currentTree.IsVisible = true
			}
			currentRow -= 1
		}
		currentCol += 1
	}

	// find number of visible trees
	for _, treeList := range t.TreeGrid {
		for _, tree := range treeList {
			if tree.IsVisible {
				numberVisibleTrees += 1
			}
		}
	}
	return numberVisibleTrees
}

func (t *TreeGridWrapper) GetScenicScore(currentRow int, currentCol int) int {

	// scenic is 0 for trees on the edge of the grid
	if currentRow == 0 || currentCol == 0 || currentRow == t.NumberOfRows-1 || currentCol == t.NumberOfCols-1 {
		return 0
	}

	viewingDistanceLeft := 0
	viewingDistanceRight := 0
	viewingDistanceUp := 0
	viewingDistanceDown := 0

	var nextRow int
	var nextCol int

	// right viewing distance
	nextCol = currentCol + 1
	for nextCol < t.NumberOfCols {
		viewingDistanceRight += 1
		if t.TreeGrid[currentRow][nextCol].Height >= t.TreeGrid[currentRow][currentCol].Height {
			break
		}
		nextCol += 1
	}

	// left viewing distance
	nextCol = currentCol - 1
	for nextCol >= 0 {
		viewingDistanceLeft += 1
		if t.TreeGrid[currentRow][nextCol].Height >= t.TreeGrid[currentRow][currentCol].Height {
			break
		}
		nextCol -= 1
	}

	// up viewing distance
	nextRow = currentRow - 1
	for nextRow >= 0 {
		viewingDistanceUp += 1
		if t.TreeGrid[nextRow][currentCol].Height >= t.TreeGrid[currentRow][currentCol].Height {
			break
		}
		nextRow -= 1
	}

	// down viewing distance
	nextRow = currentRow + 1
	for nextRow < t.NumberOfRows {
		viewingDistanceDown += 1
		if t.TreeGrid[nextRow][currentCol].Height >= t.TreeGrid[currentRow][currentCol].Height {
			break
		}
		nextRow += 1
	}
	// fmt.Printf("(up :%d, down %d, left: %d, right: %d)\t", viewingDistanceUp, viewingDistanceDown, viewingDistanceLeft, viewingDistanceRight)
	return viewingDistanceDown * viewingDistanceLeft * viewingDistanceRight * viewingDistanceUp
}

func getTreeGridWrapper(reader io.Reader) *TreeGridWrapper {
	treeGrid := make([][]*Tree, 0)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		treeList := make([]*Tree, 0)
		line := scanner.Text()
		for _, digit := range line {
			height, _ := strconv.Atoi(string(digit))
			treeList = append(treeList, &Tree{
				Height:    height,
				IsVisible: false,
			})
		}
		treeGrid = append(treeGrid, treeList)
	}

	return &TreeGridWrapper{
		TreeGrid:     treeGrid,
		NumberOfRows: len(treeGrid),
		NumberOfCols: len(treeGrid[0]),
	}
}

func main() {
	file, err := os.Open("/home/ec2-user/go/src/github.com/iamwillzhu/adventofcode2022day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	treeGridWrapper := getTreeGridWrapper(file)

	// for _, treeList := range treeGridWrapper.TreeGrid {
	// 	for _, tree := range treeList {
	// 		fmt.Printf("%d", tree.Height)
	// 	}
	// 	fmt.Println()
	// }

	numberVisibleTrees := treeGridWrapper.GetNumberVisibleTrees()

	// fmt.Println("Printing visible trees...")

	// for _, treeList := range treeGridWrapper.TreeGrid {
	// 	for _, tree := range treeList {
	// 		if tree.IsVisible {
	// 			fmt.Printf("1")
	// 		} else {
	// 			fmt.Printf("0")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	highestScenicScore := 0
	for row, treeList := range treeGridWrapper.TreeGrid {
		for col, _ := range treeList {
			scenicScore := treeGridWrapper.GetScenicScore(row, col)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
		// fmt.Println()
	}

	fmt.Printf("The number of trees visible from outside the grid is %d\n", numberVisibleTrees)
	fmt.Printf("The highest scenic score possible of a tree within the grid is %d\n", highestScenicScore)
}
