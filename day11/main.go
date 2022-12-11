package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"regexp"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type monkey struct {
	index int
	items []int
	operation func(int)int
	newMonkey func(int)int
	inspections int
}

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n\r\n")

	// ------- </preprocessing> ---------- //

	monkeys := []monkey{}
	globalModulus := 1

	for _, monke := range data {
		monkeyLines := strings.Split(monke, "\r\n")

		l := monkeyLines[3]
		n, _ := strconv.Atoi(strings.Split(l, "  Test: divisible by ")[1])
		globalModulus *= n

		newMonkey := getMonkey(monkeyLines)
		monkeys = append(monkeys, newMonkey)
	}

	rounds := 10000
	// holy for loop
	for i := 0; i < rounds; i++ {
		for mIdx := range monkeys {
			monke := monkeys[mIdx]
			for _, item := range monke.items {
				item = monke.operation(item) % globalModulus
				monke.items = monke.items[1:]
				newMonkeyIdx := monke.newMonkey(item)
				monkeys[newMonkeyIdx].items = 
				append(monkeys[newMonkeyIdx].items, item)
				monke.inspections++
			}
			monkeys[mIdx] = monke
		}
	}
	mostInspections := [2]int{0, 0}
	for _, monke := range monkeys {
		if monke.inspections >= mostInspections[0] {
			mostInspections[1] = mostInspections[0]
			mostInspections[0] = monke.inspections
		} else if monke.inspections > mostInspections[1] {
			mostInspections[1] = monke.inspections
		}
	}
	fmt.Println(mostInspections[0] * mostInspections[1])
}

func getMonkey(inp []string) monkey {
	re := regexp.MustCompile("[0-9]+")

	idx, _ := strconv.Atoi(re.FindAllString(inp[0], 1)[0])
	items := convertToInt(re.FindAllString(inp[1], -1))
	operation := parseOperation(inp[2])
	test := parseTest(inp[3:])
	return monkey {
		index: idx,
		items: items,
		operation: operation,
		newMonkey: test,
		inspections: 0,
	}
}

func emptyMonkeys(inp []monkey) []monkey {
	ret := make([]monkey, len(inp)) 
	for i, monke := range inp {
		ret[i] = monkey{
			index: monke.index,
			items: []int{},
			operation: monke.operation,
			newMonkey: monke.newMonkey,
		}
	}
	return ret
}


func convertToInt(inp []string) []int {
	ret := make([]int, len(inp))
	for i, s := range inp {
		ret[i], _ = strconv.Atoi(s)
	}
	return ret
}

func parseOperation(inp string) func(int) int {
	inp = strings.Split(inp, "  Operation: new = old ")[1]
	op := inp[0]
	num, err := strconv.Atoi(strings.Split(inp, " ")[1])
	if err != nil {
		if string(op) == "+" {
			return func(x int) int {
				return x + x
			}
		}
		return func(x int) int {
			return x * x
		}
	}
	if string(op) == "+" {
		return func(x int) int {
			return x + num
		}
	}
	return func(x int) int {
		return x * num
	}
}

func parseTest(inp []string) func(int) int {
	re := regexp.MustCompile("[0-9]+")
	test,    _ := strconv.Atoi(re.FindAllString(inp[0], 1)[0])
	ifTrue,  _ := strconv.Atoi(re.FindAllString(inp[1], 1)[0])
	ifFalse, _ := strconv.Atoi(re.FindAllString(inp[2], 1)[0])
	return func(x int) int {
		if x % test == 0 {
			return ifTrue
		}
		return ifFalse
	}
}