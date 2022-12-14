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

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")
	
	re := regexp.MustCompile("[0-9]+")
	
	// num, _ := strconv.Atoi(re.FindAllString(cmd, 1)[0])

	walls := map[complex128]int{}
	voidLimit := 0.0
	xMin, xMax := 10000.0, 0.0

	for _, line := range data {
		numsStr := re.FindAllString(line, -1)
		x1, _ := strconv.ParseFloat(numsStr[0], 64)
		y1, _ := strconv.ParseFloat(numsStr[1], 64)
		if y1 > voidLimit {
			voidLimit = y1
		}
		if x1 < xMin {
			xMin = x1
		}
		if x1 > xMax {
			xMax = x1
		}
		for i := 2; i < len(numsStr); i += 2 {
			xnext, _ := strconv.ParseFloat(numsStr[i], 64)
			ynext, _ := strconv.ParseFloat(numsStr[i+1], 64)
			
			if ynext > voidLimit {
				voidLimit = ynext
			}
			if xnext < xMin {
				xMin = xnext
			}
			if xnext > xMax {
				xMax = xnext
			}

			if xnext == x1 {
				yL := y1
				yH := ynext
				if yL > yH {
					yL, yH = yH, yL
				}
				for y := yL; y <= yH; y++ {
					walls[complex(x1, y)] = 1
				}
			} else {
				xL := x1
				xH := xnext
				if xL > xH {
					xL, xH = xH, xL
				}
				for x := xL; x <= xH; x++ {
					walls[complex(x, y1)] = 1
				}
			}
			x1, y1 = xnext, ynext
		}
	}

	// ------- </preprocessing> ---------- //

	sandsDropped := -1
	abyss := false
	for !abyss {
		walls, abyss = dropSandP1(walls, voidLimit)
		sandsDropped++
	}
	fmt.Println("Part 1:", sandsDropped)
	blocked := false
	for !blocked {
		walls, blocked = dropSandP2(walls, voidLimit)
		sandsDropped++
	}
	fmt.Println("Part 2:", sandsDropped)
}

func dropSandP1(walls map[complex128]int, voidLimit float64) (map[complex128]int, bool) {
	currentPosition := complex(500, 0)
	iters := 0
	placed := false
	for !placed {
		iters++
		if imag(currentPosition) > voidLimit {
			return walls, true
		}
		if walls[currentPosition + 1i] == 0 {
			currentPosition += 1i
			continue
		}
		if walls[currentPosition - 1 + 1i] == 0 {
			currentPosition += -1 + 1i
			continue
		}
		if walls[currentPosition + 1 + 1i] == 0 {
			currentPosition += 1 + 1i
			continue
		}
		placed = true
		walls[currentPosition] = 2
	}
	return walls, false
}

func dropSandP2(walls map[complex128]int, voidLimit float64) (map[complex128]int, bool) {
	currentPosition := complex(500, 0)
	iters := 0
	placed := false
	for !placed {
		iters++
		if imag(currentPosition) == voidLimit + 1 {
			walls[currentPosition] = 2
			return walls, false
		}
		if walls[currentPosition + 1i] == 0 {
			currentPosition += 1i
			continue
		}
		if walls[currentPosition - 1 + 1i] == 0 {
			currentPosition += -1 + 1i
			continue
		}
		if walls[currentPosition + 1 + 1i] == 0 {
			currentPosition += 1 + 1i
			continue
		}
		placed = true
		walls[currentPosition] = 2
	}
	return walls, imag(currentPosition) == 0
}

func printSand(walls map[complex128]int, voidLimit float64, xMin float64, xMax float64) {
	for y := 0; y <= int(voidLimit) + 2; y++ {
		for x := int(xMin) - 2; x <= int(xMax) + 2; x++ {
			if walls[complex(float64(x), float64(y))] == 1 {
				fmt.Print("#")
			} else if walls[complex(float64(x), float64(y))] == 2 {
				fmt.Print("o")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}