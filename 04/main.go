package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// see: https://regex101.com/r/VRbjzk/1
	var re = regexp.MustCompile(`.(.*)(]\s)(.[a-zA-Z]*)(..)(.*)`)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Strings(lines)
	currentGuard := 0
	guardSleepTimes := make(map[int]*[61]int)
	fmt.Println(time.Since(start))
	sleepTimer := 0 // time guard fell asleep
	for _, line := range lines {
		eventLog := re.FindAllStringSubmatch(line, -1)
		minuteSplit := strings.Split(eventLog[0][1], ":")
		minutes, _ := strconv.Atoi(minuteSplit[1])
		logType := eventLog[0][3] // falls,wakes,Guard
		if logType == "Guard" {
			s := strings.Split(eventLog[0][5], " ")
			currentGuard, err = strconv.Atoi(s[0])
			if err != nil {
				fmt.Println(err)
			}
			if guardSleepTimes[currentGuard] == nil {
				guardSleepTimes[currentGuard] = &[61]int{}
			}

		}
		if logType == "falls" {
			sleepTimer = minutes
		}
		// when guard wakes up calculate time sleep
		if logType == "wakes" {
			for i := sleepTimer; i < minutes; i++ {
				guardSleepTimes[currentGuard][i]++  // incrememnt current slept count
				guardSleepTimes[currentGuard][60]++ // total minutes slept buffer
			}
			sleepTimer = 0

		}

	}
	fmt.Println(time.Since(start))
	// Find sleepiest Guard
	sleepiestGuard := 0
	sleepiestMinute := 0
	sleepBuf := 0
	sleepiestGuardMin := 0
	for guard, sleepTime := range guardSleepTimes {
		// only get guards that have slept
		if sleepiestGuard == 0 {
			sleepiestGuard = guard

		}
		if sleepTime[60] == 0 {
			continue
		}
		if sleepTime[60] >= guardSleepTimes[sleepiestGuard][60] {
			sleepiestGuard = guard

		}
		for minute, sleepMin := range sleepTime {
			if sleepMin > sleepBuf && minute != 60 {
				sleepiestMinute = minute
				sleepBuf = sleepMin
				sleepiestGuardMin = guard
			}
		}

	}
	fmt.Println(time.Since(start))
	// Find sleepiest hour for sleepiest guard
	sleepiestHour := 0
	sleepBuf = 0
	for hour, sleepTime := range guardSleepTimes[sleepiestGuard] {
		if sleepTime > sleepBuf && hour != 60 {
			sleepiestHour = hour
			sleepBuf = sleepTime
		}
	}
	fmt.Println("Sleepiest Guard:", sleepiestGuard)
	fmt.Println("Sleepiest hour:", sleepiestHour)
	fmt.Println("Part 1 Code:", sleepiestGuard*sleepiestHour)
	fmt.Println("Part 2 Code:", sleepiestGuardMin*sleepiestMinute)
	fmt.Println(time.Since(start))
}
