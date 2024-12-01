package main

import (
	"bufio"
	"log"
	"os"
	"sort"
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

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, "   ")
		if len(numbers) < 2 {
			log.Fatal("length is shorter than expected 2")
		}

		intLeft, err := strconv.Atoi(numbers[0])

		if err != nil {
			log.Fatal(err)
		}

		intRight, err := strconv.Atoi(numbers[1])

		if err != nil {
			log.Fatal(err)
		}

		left = append(left, intLeft)
		right = append(right, intRight)
	}

	log.Printf("total Left: %v,Right: %v", len(left), len(right))

	sort.Ints(left)
	sort.Ints(right)

	if len(left) != len(right) {
		log.Fatal("Length of arrays does not match")
	}

	total := 0

	for ind, valLeft := range left {
		valRight := right[ind]

		diff := valLeft - valRight
		if diff < 0 {
			diff = -diff
		}

		log.Printf("L: %v R: %v D: %v", valLeft, valRight, diff)

		total += diff
	}

	log.Printf("Total Difference: %v", total)
}
