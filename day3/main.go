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

	d, err := os.ReadFile("data.txt")
	check(err)

	data := strings.Split(string(d), "\r\n")

	// ------- </preprocessing> ---------- //

	indices := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	sum := 0
	for _, v := range data {
		p1, p2 := v[:len(v)/2], v[len(v)/2:]
		for _, ss := range p1 {
			s := string(ss)
			if strings.Contains(p2, s) {
				sum += strings.Index(indices, s) + 1
				break
			}
		}
	}
	fmt.Println("Part 1:", sum)

	sum = 0
	chunkSize := 3
	for i := 0; i < len(data)/chunkSize; i++ {
		chunk := data[3*i:3*i+chunkSize]
		for _, ss := range chunk[0] {
			s := string(ss)
			if strings.Contains(chunk[1], s) && strings.Contains(chunk[2], s) {
				sum += strings.Index(indices, s) + 1
				break
			}
		}
	}
	fmt.Println("Part 2:", sum)
}

