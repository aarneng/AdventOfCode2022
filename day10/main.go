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

func addSignalStrength(cycle int, sum int, spritePos int, num int) (int, int) {
	currentPos := (cycle % 40) - 2
	if currentPos - 1 == spritePos || currentPos == spritePos || currentPos + 1 == spritePos {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if (cycle % 40) == 1 {
		fmt.Println()
	}
	spritePos += num
	if cycle % 40 == 20 && cycle <= 220 {
		return sum + spritePos * cycle, spritePos
	}
	return sum, spritePos
}


func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	x := 1
	sum := 0
	cycle := 1

	for _, instruction := range data {
		cycle++
		sum, x = addSignalStrength(cycle, sum, x, 0)

		if instruction == "noop" {
			continue
		}

		num, _ := strconv.Atoi(strings.Split(instruction, " ")[1])

		cycle++
		sum, x = addSignalStrength(cycle, sum, x, num)
	}
	fmt.Println(sum)
}

