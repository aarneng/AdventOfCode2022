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


func isBetween(a1 int, a2 int, b1 int, b2 int) bool {
	return b1 >= a1 && b2 <= a2
}

func isOverlap(a1 int, a2 int, b1 int, b2 int) bool {
	return (a1 <= b1 && b1 <= a2) || (a1 <= b2 && b2 <= a2)
}


func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	count1 := 0
	count2 := 0

	for _, vv := range data {
		v := strings.Split(vv, ",")
		v1 := strings.Split(v[0], "-")
		v2 := strings.Split(v[1], "-")
		i1, err := strconv.Atoi(v1[0])
		check(err)
		i2, err := strconv.Atoi(v1[1])
		check(err)
		j1, err := strconv.Atoi(v2[0])
		check(err)
		j2, err := strconv.Atoi(v2[1])
		check(err)

		if isBetween(i1, i2, j1, j2) || isBetween(j1, j2, i1, i2) {
			count1++
			count2++
		} else if isOverlap(i1, i2, j1, j2) {
			count2++
		}
	}

	fmt.Println("Part 1:", count1)
	fmt.Println("Part 2:", count2)
}

