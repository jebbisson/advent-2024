package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	total := 0

	completeText := ""

	for scanner.Scan() {
		text := scanner.Text()
		completeText += text
	}

	//split on dos
	dos := strings.Split(completeText, "do()")

	//reset so we can only regex valid do statements
	completeText = ""

	for _, val := range dos {
		//find any don't in each do entry
		donts := strings.Split(val, "don't()")
		//since there are no do's in each val, dont means ignore all remaining.
		if len(donts) > 0 {
			//add only before the donts
			completeText += donts[0]
		} else {
			//if not donts are found then block is good
			completeText += val
		}
	}

	log.Printf("T: %v", completeText)

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := r.FindAllString(completeText, -1)
	if matches == nil {
		log.Print("no matches found")
	}

	for _, val := range matches {
		val = strings.ReplaceAll(val, "mul(", "")
		val = strings.ReplaceAll(val, ")", "")

		numbers := strings.Split(val, ",")

		if len(numbers) < 2 {
			log.Fatal("Less than 2 numbers found")
		}

		leftNum, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}

		rightNum, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("L: %v R: %v", leftNum, rightNum)

		total += leftNum * rightNum

	}

	log.Printf("Total: %v", total)

	//184576302
	//302749809
}
