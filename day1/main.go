package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	temp := strings.Split(string(d), "\r\n\r\n")
	data := make([][]int, len(temp))
	for i, s := range temp {
		tmp := strings.Fields(s)
		data[i] = strArrToIntArr(tmp)
	}

	// ------- </preprocessing> ---------- //

	// data is an array of array of ints

	maxVals := make([]int, 3)
	idxSmallest, valSmallest := 0, -1

	for _, arr := range data {
		v := arraySum(arr)
		if v > valSmallest {
			maxVals[idxSmallest] = v
			idxSmallest, valSmallest = idxValByFunc(maxVals, func(a int, b int) int {return b - a})
		}
	}

	// ------ part 1 -------- //
	maxIdx, _ := idxValByFunc(maxVals, func(a int, b int) int {return a - b})
	maxVal := maxVals[maxIdx]
	fmt.Println("Most Calories carried by single elf:", maxVal)


	// ------ part 2 -------- //
	sumOfMax := arraySum(maxVals)
	fmt.Println("Calories carried by top 3 elves:", sumOfMax)
}

type maximiser func(int, int) int

func idxValByFunc(arr []int, fn maximiser) (int, int) {
	idx, val := 0, arr[0]
	for i, v := range arr[0:] {
		if fn(val, v) < 0 {
			idx = i
			val = v
		}
	}
	return idx, val
}

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func strArrToIntArr(strs []string) []int {
	var ret = make([]int, len(strs))
	for i, s := range strs {
		j, err := strconv.Atoi(s)
		check(err)
		ret[i] = j
	}
	return ret
}
