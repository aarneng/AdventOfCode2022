package main

import (
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}


func abs(n int) int {
	// this is part of the math package,
	// but that implementation returns a float
	if n > 0 {
		return n
	}
	return -n
}

func distance(head [2]int, tail [2]int) int {
	if (head[0] == tail[0]) && (head[1] == tail[1]) {
		return 0
	}
	if abs(head[0] - tail[0]) + abs(head[1] - tail[1]) == 1 || 
		 (abs(head[0] - tail[0]) == 1 && abs(head[1] - tail[1]) == 1) {
		return 1
	}
	return 2
}


func walkTail(head [2]int, tail [2]int) [2]int {
	// walk the "tail" knot towards the "head" knot

	if distance(head, tail) <= 1 {
		// don't move
	} else if (head[0] == tail[0]) {
		// walk along y axis
		d2 := head[1] - tail[1]
		d2 = d2 / abs(d2)
		tail[1] += d2
	} else if (head[1] == tail[1]) {
		// walk along x axis
		d2 := head[0] - tail[0]
		d2 = d2 / abs(d2)
		tail[0] += d2
	} else {
		// walk diagonally
		d1 := head[0] - tail[0]
		d1 = d1 / abs(d1)
		d2 := head[1] - tail[1]
		d2 = d2 / abs(d2)
		tail[0] += d1
		tail[1] += d2
	}
	return tail
}


func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	const ropeLength = 10 // <---- change this to 2 for part 1 and 10 for part 2
	currPoints := [ropeLength][2]int{}
	uniquePoints := map[[2]int]bool{}

	for _, input := range data {
		s := strings.Split(input, " ")
		d, a := s[0], s[1]
		amt, _ := strconv.Atoi(a)

		var direction [2]int
		if d == "U" {
			direction = [2]int{0, 1}
		}		
		if d == "R" {
			direction = [2]int{1, 0}
		}
		if d == "D" {
			direction = [2]int{0, -1}
		}
		if d == "L" {
			direction = [2]int{-1, 0}
		}

		// move the head and it's tails `amt` times
		for a := 0; a < amt; a++ {

			// initialise the head by moving it by one unit
			currPoints[0][0] += direction[0]
			currPoints[0][1] += direction[1]

			for i := 0; i < ropeLength - 1; i++ {
				// make the rest of the knots follow the one in front of them
				currPoints[i+1] = walkTail(currPoints[i], currPoints[i+1])
			}

			// mark the position of the last knot as visited
			uniquePoints[currPoints[ropeLength-1]] = true
		}
	}
	fmt.Println(len(uniquePoints))
}

