package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
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

	sizes := map[string]int{"/": 0}
	dirPath := []string{"/"}
	
	re := regexp.MustCompile("[0-9]+")

	for _, cmd := range data {
		if cmd == "$ cd /" {
			// go to root dir
			dirPath = []string{"/"}
			continue
		}
		if cmd == "$ cd .." {
			// go back one level, aka remove last item from path
			dirPath = dirPath[:len(dirPath) - 1]
			continue
		}
		if cmd[:4] == "$ cd" {
			// check if we change directory to somewhere other than "/" or ".."
			// if so, save the name of that dir to the path
			path := strings.Split(cmd, " ")[2]
			dirPath = append(dirPath, path)
			continue
		}
		if (cmd[:3] == "dir") || (cmd == "$ ls") {
			// we con't need to process anything from these commands
			continue
		}
		// if none of the if statements above get triggered,
		// it means that we are currently reading the size of a file
		num, _ := strconv.Atoi(re.FindAllString(cmd, 1)[0])

		// create a variable to keep track of the path
		// we will append to it later. This is needed as
		// there can be duplicate directory nested within one another,
		// so if we were to save to the (key, value) pair
		// corresponding to (dirName, fileSize) we'd get collisions.
		// Saving as (pathToDir, fileSize) instead avoids the collision.
		currDirPath := "/"
		sizes[currDirPath] += num
		
		for _, dir := range dirPath[1:] {
			currDirPath += string(dir) + "/"
			sizes[currDirPath] += num
		}
	}

	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}
	
	fmt.Println("part 1:", sum)

	minSize := sizes["/"] - 40000000
	currentMin := sizes["/"]
	for _, size := range sizes {
		if size > minSize && size < currentMin {
			currentMin = size
		}
	}

	fmt.Println("part 2:", currentMin)
}

