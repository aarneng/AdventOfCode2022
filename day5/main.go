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

func getBoxIdx(i int) int {
	/*
		return the index where the value of the input is stored
		for example, if a row has values
		"[S] [B] [B] [F] [H] [C] [B] [N] [L]"
		  1   2   3   4   5   6   7   8   9
		then we can get the values by looping through indices 0 - 8 
		and index 0 corresponds to value S, found at string index 1, 
		index 1 to value B at string index 4, etc...
	*/
	return 4 * i + 1
}

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	amtBoxes := 9

	boxes := make([][]string, amtBoxes)

	for _, v := range data[:amtBoxes] {
		for i := 0; i < amtBoxes; i++ {
			j := getBoxIdx(i)
			if string(v[j]) == " " {
				continue
			}
			boxes[i] = append(boxes[i], string(v[j]))
		}
	}

	re := regexp.MustCompile("[0-9]+")

	for _, v := range data[amtBoxes+1:] {
		vals := re.FindAllString(v, 3)
		howMany, _:= strconv.Atoi(vals[0])
		from, _ := strconv.Atoi(vals[1])
		to, _ := strconv.Atoi(vals[2])

		from, to = from - 1, to - 1

		// -------- part 1 ------------
		
		//* <------ toggle this comment to/from //* and /* to toggle parts 1 & 2
		for i := 0; i < howMany; i++ {
			item := boxes[from][0]
			boxes[from] = boxes[from][1:]
			boxes[to] = append([]string{item}, boxes[to]...)
		}
		/*/ //-------- part 2 ------------
		items := make([]string, howMany)
		copy(items, boxes[from][0:howMany])
		boxes[from] = boxes[from][howMany:]
		boxes[to] = append(items, boxes[to]...)
		//*/
	}

	ans := ""
	for _, v := range boxes {
		ans += v[0]
	}
	fmt.Println(ans)
}

