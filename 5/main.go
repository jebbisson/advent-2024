package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("rules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	ruleList := make(map[int]map[int]bool)

	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, "|")

		if len(slice) < 2 {
			log.Fatal("text string must be longer than 2")
		}

		numleft, err := strconv.Atoi(slice[0])

		if err != nil {
			log.Fatal(err)
		}

		numright, err := strconv.Atoi(slice[1])

		if err != nil {
			log.Fatal(err)
		}

		rule, ok := ruleList[numleft]
		if ok {
			rule[numright] = true
		} else {
			ruleList[numleft] = map[int]bool{numright: true}
		}
	}

	file2, err2 := os.Open("pages.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer func() {
		if err2 = file2.Close(); err2 != nil {
			log.Fatal(err2)
		}
	}()

	scanner2 := bufio.NewScanner(file2)

	numGrid := [][]int{}

	for scanner2.Scan() {
		text := scanner2.Text()
		slice := strings.Split(text, ",")

		ints := make([]int, len(slice))

		for i, val := range slice {
			v, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			ints[i] = v
		}

		numGrid = append(numGrid, ints)
		//completeText += text
	}

	successGrid := [][]int{}

	for _, line := range numGrid {
		success := true
		length := len(line)
		for i, num := range line {
			if i+1 == length {
				break
			}
			//!!!need to fix boundary limit issues!!!
			test := ruleList[line[i+1]][num]
			if test { // if we find an opposing rule then we are done
				success = false
				break
			}
		}

		if success {
			successGrid = append(successGrid, line)
		}
	}

	total := 0
	for _, val := range successGrid {
		//mid := val[math.Floor(len(val)/2)]
		mid := len(val) / 2
		log.Printf("Success line: %v, mid: %v", val, val[mid])
		total += val[mid]
	}

	log.Printf("Total mid %v", total)
}
