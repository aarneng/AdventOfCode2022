package main

import (
	"fmt"
	"os"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkAllUnique(m map[rune]int) bool {
	for _, v := range m {
		if v >= 2 {
			return false
		}
	}
	return true
}

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := []rune(string(d))

	// ------- </preprocessing> ---------- //

	windowLength := 14    // <--- change this to 4 if you want part 1, 14 for part 2
	pointValues := map[rune]int{}
	
	for _, c := range data[:windowLength] {
		pointValues[c]++
	}

	if checkAllUnique(pointValues) {
		// this check probably isn't needed but it's here for completion's sake
		fmt.Println(windowLength)
		return
	}

	for i, c := range data[windowLength:] {
		pointValues[data[i]]--
		pointValues[c]++
		if checkAllUnique(pointValues) {
			fmt.Println(i + windowLength + 1)
			break
		}
	}
}

