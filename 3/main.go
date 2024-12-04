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

	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("T: %v", text)

		matches := r.FindAllString(text, -1)
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
	}

	log.Printf("Total: %v", total)
}
