package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
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
	myImage := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	// build a map of coors with claims
	for _, c := range claims {
		for y := c.left; y < c.left+c.width; y++ {
			for x := c.top; x < c.top+c.height; x++ {
				col := color.RGBA{255, 0, 0, 100}
				myImage.Set(x, y, col)
				// if overlapping, set color hue darker for number of overlaps
				if grid[x][y] == 1 {
					col = color.RGBA{255, 0, 0, 100 + uint8(grid[x][y]*20)}
					myImage.Set(x, y, col)
					overlapCounter++
				}
				grid[x][y]++
			}
		}
	}
	// find non overlapping claim
	nonOverlap := 0
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
			nonOverlap = c.claimId
			fmt.Printf("Non overlapping claim: %v\n", c.claimId)
		}
	}
	// make non overlapping square green
	for _, c := range claims {
		if c.claimId == nonOverlap {
			for y := c.left; y < c.left+c.width; y++ {
				for x := c.top; x < c.top+c.height; x++ {
					col := color.RGBA{0, 255, 0, 255}
					myImage.Set(x, y, col)
				}
			}
		}
	}
	outputFile, _ := os.Create("test.png")
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	outputFile.Close()
	fmt.Println("Overlapping inches:", overlapCounter)
	fmt.Println(time.Since(start))
}
