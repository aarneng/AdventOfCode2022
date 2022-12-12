package main

import (
	"fmt"
	"os"
	_ "strconv"
	"strings"
	_ "regexp"
	_ "reflect"
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

	chunkSize := len(strings.Split(string(d), "\r\n")[0])
	chunks := [][]uint8{}
	startPos := [2]int{}
	endPos := [2]int{}
	for i := 0; i < len(d); i += chunkSize + 2{
    end := i + chunkSize

		i1 := indexOf(d[i:end], 83)
		if i1 != -1 {
			startPos = [2]int{i/chunkSize, i1}
		}		
		i2 := indexOf(d[i:end], 69) // nice
		if i2 != -1 {
			endPos = [2]int{i/chunkSize, i2}
		}

    if end > len(d) {
        end = len(d)
    }

    chunks = append(chunks, d[i:end])
	}

	chunks[startPos[0]][startPos[1]] = 97
	chunks[endPos[0]][endPos[1]] = 123

	// ------- </preprocessing> ---------- //

	visited := map[[2]int]bool{startPos: true}
	part1 := 0
	for !visited[endPos] {
		visited = allNearby(visited, chunks, chunkSize, len(d) / chunkSize)
		part1++
	}
	fmt.Println("part 1:", part1)

	shortest := part1
	for i := 0; i < len(chunks[0]); i++ {
		for j := 0; j < len(chunks); j++ {
			k := chunks[j]
			if k[i] == 97 {
				startPos = [2]int{j, i}
				visited := map[[2]int]bool{startPos: true}
				runs := 0
				for !visited[endPos] && runs < shortest {
					visited = allNearby(visited, chunks, chunkSize, len(d) / chunkSize)
					runs++
				}
				if runs < shortest {
					shortest = runs
				}
			}
		}
	}
	fmt.Println(shortest)
}


func allNearby(allPos map[[2]int]bool, chunks [][]uint8, maxW int, maxH int) map[[2]int]bool {
	for _, k := range getKeys(allPos) {
		for _, nearby := range nearbyOneSquare(k, chunks, maxW, maxH) {
			allPos[nearby] = true
		} 
	}
	return allPos
}


func nearbyOneSquare(pos [2]int, chunks [][]uint8, maxW int, maxH int) [][2]int {
	ret := [][2]int{}
	if pos[0] > 0 {
		newPos := [2]int{pos[0] - 1, pos[1]}
		if chunks[newPos[0]][newPos[1]] <= chunks[pos[0]][pos[1]] + 1 {
			ret = append(ret, newPos)
		}
	}
	if pos[1] > 0 {
		newPos := [2]int{pos[0], pos[1] - 1}
		if chunks[newPos[0]][newPos[1]] <= chunks[pos[0]][pos[1]] + 1 {
			ret = append(ret, newPos)
		}
	}
	if pos[0] < maxH - 2 {
		newPos := [2]int{pos[0] + 1, pos[1]}
		cv2 := chunks[pos[0]][pos[1]]
		cv1 := chunks[newPos[0]][newPos[1]]
		if cv1 <= cv2 + 1 {
			ret = append(ret, newPos)
		}
	}
	if pos[1] < maxW - 1 {
		newPos := [2]int{pos[0], pos[1] + 1}
		if chunks[newPos[0]][newPos[1]] <= chunks[pos[0]][pos[1]] + 1 {
			ret = append(ret, newPos)
		}
	}
	return ret
}

func getKeys(myMap map[[2]int]bool) [][2]int {
	keys := make([][2]int, len(myMap))

	i := 0
	for k := range myMap {
			keys[i] = k
			i++
	}

	return keys
}


func indexOf(nums []uint8, n int) int {
	num := uint8(n)
	for i, u := range nums {
		if u == num {
			return i
		}
	}
	return -1
}