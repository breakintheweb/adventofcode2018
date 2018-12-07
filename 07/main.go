package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func processStep(steps map[rune][]rune, currentStep rune) {
	// for all steps remianing
	for stepIndex, stepRune := range steps {
		ts := []rune{}
		// for each step
		for _, parentStep := range stepRune {
			// if not equal to the rune we are deleting
			if parentStep != currentStep {
				ts = append(ts, parentStep)
			}
		}
		steps[stepIndex] = ts
	}
	delete(steps, currentStep)
}
func getNextStepR(sb map[rune][]rune) rune {
	var v rune
	readySteps := []rune{}
	// get list of steps without pending parent steps
	for f, e := range sb {
		if len(e) == 0 {
			readySteps = append(readySteps, f)
		}
	}
	// sort or next step buffer alphabetically
	sort.Slice(readySteps, func(i, j int) bool {
		return readySteps[i] < readySteps[j]
	})

	for _, readyStep := range readySteps {
		processStep(sb, readyStep)
		return readyStep
	}
	// should never reach here
	return v
}

func getNextStep(sb map[rune][]rune, workers map[rune]int) {
	readySteps := []rune{}
	// get list of steps without pending parent steps
	for f, e := range sb {
		if len(e) == 0 {
			readySteps = append(readySteps, f)
		}
	}
	// sort our next step buffer alphabetically
	sort.Slice(readySteps, func(i, j int) bool {
		return readySteps[i] < readySteps[j]
	})
	for _, readyStep := range readySteps {
		if _, ok := workers[readyStep]; ok {
		} else if len(workers) <= 5 {
			workers[readyStep] = int(readyStep - 4)
		}
	}
}

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	steps := make(map[rune][]rune)
	steps1 := make(map[rune][]rune)

	scanner := bufio.NewScanner(file)
	workers := make(map[rune]int, 0)
	for scanner.Scan() {
		var sc rune
		var sn rune
		fmt.Sscanf(scanner.Text(), "Step %c must be finished before step %c can begin.", &sc, &sn)
		if _, ok := steps[sc]; !ok {
			steps[sc] = []rune{}
			steps1[sc] = []rune{}
		}
		if _, ok := steps[sn]; !ok {
			steps[sn] = []rune{}
			steps1[sn] = []rune{}

		}
		steps[sn] = append(steps[sn], sc)
		steps1[sn] = append(steps1[sn], sc)

	}
	stepOrder := []rune{}
	timer := 0
	// Part 1
	for i := len(steps1); i > 0; {
		nextStep := getNextStepR(steps1)
		stepOrder = append(stepOrder, nextStep)
		i--
	}
	// Part 2
	for i := len(steps); i > 0; {
		workerstmp := make(map[rune]int, 0)
		for iw, v := range workers {
			if workers[iw] != 0 {
				workerstmp[iw] = v - 1
			} else {
				processStep(steps, iw)
				i--
			}
		}
		workers = workerstmp
		// process queue
		getNextStep(steps, workers)
		timer++
	}
	fmt.Printf("Part 1: %s\n", string(stepOrder))
	fmt.Printf("Part 2: %v\n", timer)
	fmt.Println(time.Since(start))

}
