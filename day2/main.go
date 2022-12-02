package main

import (
	"fmt"
	"os"
	"strings"
)
	
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// ------- <preprocessing> ---------- //

	d, err := os.Open("data.txt")
	check(err)

	defer d.Close()

	fmt.Println(d)

	// data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	myPoints1 := 0
	myPoints2 := 0
	var pointValues = map[string]int {
		"A":1, "B":2, "C":3,
		"X":1, "Y":2, "Z":3,
	}
	for _, v := range data {
		opp := pointValues[string(v[0])]
		my := pointValues[string(v[2])]

		// ---- part 1 ------

		myPoints1 += my
		if (opp - my + 3) % 3 == 0 {
			// draw
			myPoints1 += 3
		}
		if (opp - my + 3) % 3 == 2 {
			// win
			myPoints1 += 6
		}

		// ---- part 2 ------

		if string(v[2]) == "X" {
			if opp == 1 {
				myPoints2 += 3
			} else {
				myPoints2 += opp - 1
			}
		}
		if string(v[2]) == "Y" {
			myPoints2 += opp + 3
		}
		if string(v[2]) == "Z" {
			if opp == 2 {
				myPoints2 += 9
			} else {
				myPoints2 += (opp + 1) % 3 + 6
			}
		}
		
	}
	fmt.Println(myPoints1)
	fmt.Println(myPoints2)
}

