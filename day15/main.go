package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type coords struct {
	sx int
	sy int
	bx int
	by int
}


func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")
	
	re := regexp.MustCompile(`[-]{0,1}[\d]*[.]{0,1}[\d]+`)

	coordsArray := []coords{}
	
	for _, line := range data {
		n := re.FindAllString(line, 4)
		newCoords := coords{toInt(n[0]), toInt(n[1]), toInt(n[2]), toInt(n[3])}
		coordsArray = append(coordsArray, newCoords)
	}

	// ------- </preprocessing> ---------- //
	
	filledSquares := map[int]bool{}
	yLine := 2000000
	for _, c := range coordsArray {
		distance := abs(c.sx - c.bx) + abs(c.sy - c.by)
		if distance < abs(c.sy - yLine) {
			continue
		}
		distance -= abs(c.sy - yLine)
		for i := c.sx - distance; i <= c.sx + distance; i++ {
			filledSquares[i] = true
		}
	}
	for _, c := range coordsArray {
		if c.by == yLine {
			delete(filledSquares, c.bx)
		}
	}
	fmt.Println("part 1:", len(filledSquares))

	availableSquares := map[int]rangeValues{}
	maxVal1 := 4000000
	maxVal2 := 4000000
	for _, c := range coordsArray {
		d := abs(c.sx - c.bx) + abs(c.sy - c.by)
		i := 0
		for x := c.sx - d; x <= c.sx + d; x++ {
			if 0 <= x && x <= maxVal1 {
				yRange := [2]int{max(c.sy - i, 0), min(c.sy + i, maxVal1)}
				a := availableSquares[x]
				a.insert(yRange)
				availableSquares[x] = a
			}
			if x >= c.sx {	
				i--
			} else {
				i++
			}
		}
	}

	for k, e := range availableSquares {
		if len(e.xVals) != 1 {
			fmt.Println("Part 2:", k * maxVal2 + e.xVals[0][1] + 1)
		}
	}
}

type rangeValues struct {
	xVals [][2]int
}

func correctIdx(arr [][2]int, itm [2]int) int {
	for i := len(arr)-1; i >= 0; i-- {
		a := arr[i]
		if itm[0] == a[0] && itm[1] > a[1] {
			return i + 1
		}
		if itm[0] > a[0] {
			return i + 1
		}
	}
	return 0
}

func (r *rangeValues) insert(n [2]int) {
	idx := correctIdx(r.xVals, n)
	if len(r.xVals) == 0 {
		r.xVals = [][2]int{n}
	} else if len(r.xVals) == idx {
		r.xVals = append(r.xVals, n)
	} else {
		r.xVals = append(r.xVals[:idx+1], r.xVals[idx:]...)
		r.xVals[idx] = n
	}
	newVals := [][2]int{r.xVals[0]}
	for _, i := range r.xVals[1:] {
		if newVals[len(newVals)-1][1] >= i[0] - 1 {
			newVals[len(newVals)-1][1] = max(newVals[len(newVals)-1][1], i[1])
		} else {
			newVals = append(newVals, i)
		}
	}
	r.xVals = newVals
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}


func abs(a int) int {
	if a < 0 {
		return - a
	}
	return a
}

func toInt(s string) int {
	ans, e := strconv.Atoi(s)
	check(e)
	return ans
}