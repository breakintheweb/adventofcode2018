package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Claim struct {
	claimId, top, left, width, height int
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Make and init grid of slices
	grid := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]byte, 1000)
	}
	// Get just numeric values w/length between 1 and 4 as groups
	// Example:: #1 @ 185,501: 17x15
	// Return :: 1 185 501  17 15
	var re = regexp.MustCompile("[0-9]{1,4}")

	overlapCounter := 0
	var claims []*Claim
	for scanner.Scan() {
		claim := re.FindAllStringSubmatch(scanner.Text(), -1)
		claimId, _ := strconv.Atoi(claim[0][0])
		top, _ := strconv.Atoi(claim[2][0])
		left, _ := strconv.Atoi(claim[1][0])
		width, _ := strconv.Atoi(claim[3][0])
		height, _ := strconv.Atoi(claim[4][0])
		claims = append(claims, &Claim{claimId: claimId, top: top, left: left, width: width, height: height})
	}
	// build a map of coors with claims
	for _, c := range claims {
		for y := c.left; y < c.left+c.width; y++ {
			for x := c.top; x < c.top+c.height; x++ {
				if grid[x][y] == 1 {
					overlapCounter++
				}
				grid[x][y]++
			}
		}
	}
	// find non overlapping claim
	for _, c := range claims {
		isOverlap := false
		for y := c.left; y < c.left+c.width; y++ {
			for x := c.top; x < c.top+c.height; x++ {
				if grid[x][y] > 1 {
					isOverlap = true
				}
			}
		}
		if !isOverlap {
			fmt.Printf("Non overlapping claim: %v\n", c.claimId)
		}
	}
	fmt.Println("Overlapping inches:", overlapCounter)
	fmt.Println(time.Since(start))
}
