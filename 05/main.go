package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

var React = make([]int, 26)

func isUpper(s rune) (rune, bool) {
	su := unicode.ToUpper(s)
	if !unicode.IsUpper(s) {
		return su, false
	}
	return su, true
}
func react(s string, ltr rune) int {
	filter := func(r rune) rune {
		if strings.IndexRune(string(ltr), r) < 0 {
			return r
		}
		return -1
	}
	filter2 := func(r rune) rune {
		if strings.IndexRune(string(unicode.ToUpper(ltr)), r) < 0 {
			return r
		}
		return -1
	}

	s = strings.Map(filter, s)
	s = strings.Map(filter2, s)

	var x = make([]rune, len(s))
	var y = make([]rune, len(s))
	var z = make([]bool, len(s))
	for a, val := range s {
		x[a] = val
		y[a], z[a] = isUpper(val)
	}
	for {
		l := len(x)
		for i := 1; i < l; i++ { // we start at one since we are comparing two values
			if y[i] == y[i-1] && z[i] != z[i-1] {
				// delete current and previous char from slice
				x = append(x[:i-1], x[i+1:]...)
				y = append(y[:i-1], y[i+1:]...)
				z = append(z[:i-1], z[i+1:]...)
				break
			}
		}
		if l == len(x) {
			break
		}
	}

	return len(x)
}

const alpha = "abcdefghijklmnopqrstuvwxyz"

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileStr := ""
	for scanner.Scan() {
		fileStr = scanner.Text()
	}
	// pass dummy rune for pass 1
	p := react(fileStr, rune('\n'))

	for i, char := range alpha {
		React[i] = react(fileStr, char)
		fmt.Printf("Pass %v of %v complete\n", i+1, len(alpha))
	}
	h := 999999 // should change to know max value
	ll := ""
	for i, char := range alpha {
		if React[i] < h {
			h = React[i]
			ll = string(char)
		}
	}
	fmt.Printf("Part 1: %v\n", p)
	fmt.Printf("Part 2: %v for letter: %s\n", h, ll)
	fmt.Println(time.Since(start))

}
