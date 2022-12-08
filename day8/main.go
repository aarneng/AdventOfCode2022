package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	visibleTrees := map[string]bool{}


	// --------- part 1 ------------- //

	// the following 4 loops could probably be changed
	// to instead rotate the data 4x and have a "single"
	// loop though the rotated data
	for i, substr := range data {
		maxHeight := -1
		for j, h := range substr {
			// looking from the top
			str := fmt.Sprintf("%d,%d",i, j)
			height, _ := strconv.Atoi(string(h))
			if height > maxHeight {
				maxHeight = height
				visibleTrees[str] = true
			}
			if maxHeight == 9 {
				break
			}
		}
		maxHeight = -1
		for j := len(substr) - 1; j >= 0; j--  {
			// looking from the bottom
			str := fmt.Sprintf("%d,%d",i, j)
			height, _ := strconv.Atoi(string(substr[j]))
			if height > maxHeight {
				maxHeight = height
				visibleTrees[str] = true
			}
			if maxHeight == 9 {
				break
			}
		}
	}
	for j := range data[0] {
		maxHeight := -1
		for i := range data {
			// looking from the left
			str := fmt.Sprintf("%d,%d",i, j)
			height, _ := strconv.Atoi(string(data[i][j]))
			if height > maxHeight {
				maxHeight = height
				visibleTrees[str] = true
			}
			if maxHeight == 9 {
				break
			}
		}
		maxHeight = -1
		for i := len(data) - 1; i >= 0; i--  {
			// looking from the right
			str := fmt.Sprintf("%d,%d",i, j)
			height, _ := strconv.Atoi(string(data[i][j]))
			if height > maxHeight {
				maxHeight = height
				visibleTrees[str] = true
			}
			if maxHeight == 9 {
				break
			}
		}
	}

	fmt.Println("part 1:", len(visibleTrees))


	// --------- part 2 ------------- //

	// really ugly code. Essentially the logic is as follows:
	// go through each tree not on the edges, and count how many trees 
	// you see in each direction (N, E, S, W), then  multiply 
	// scenic score by amt_trees
	// I have no idea if there's a cleaner or more efficient
	// solution for this
	maxScenicScore := 0
	for ii, substr := range data {
		if ii == 0 || ii == len(data) - 1 {
			continue
		}
		for jj, h := range substr {
			if jj == 0 || jj == len(substr) - 1 {
				continue
			}
			height, _ := strconv.Atoi(string(h))
			i := ii - 1
			currScenicScore := 1
			for (i >= 0) {
				// from current tree, look up
				otherHeight, _ := strconv.Atoi(string(data[i][jj]))
				if otherHeight >= height || i == 0 {
					currScenicScore *= (ii - i)
					break
				}
				i--
			}
			i = ii + 1
			for (i <= len(data) - 1) {
				// from current tree, look down
				otherHeight, _ := strconv.Atoi(string(data[i][jj]))
				if otherHeight >= height || i == len(data) - 1 {
					currScenicScore *= (i - ii)
					break
				}
				i++
			}
			j := jj - 1
			for (j >= 0) {
				// fropm current tree, look left
				otherHeight, _ := strconv.Atoi(string(data[ii][j]))
				if otherHeight >= height || j == 0 {
					currScenicScore *= (jj - j)
					break
				}
				j--
			}
			j = jj + 1
			for (j <= len(substr) - 1) {
				// from current tree, look right
				otherHeight, _ := strconv.Atoi(string(data[ii][j]))
				if otherHeight >= height || j == len(substr) - 1 {
					currScenicScore *= (j - jj)
					break
				}
				j++
			}
			if currScenicScore > maxScenicScore {
				maxScenicScore = currScenicScore
			}
		}
	}
	fmt.Println("Part 2:", maxScenicScore)
}

