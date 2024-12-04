package main

import (
	"bufio"
	"log"
	"os"
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

	successCount := 0

	for scanner.Scan() {
		text := scanner.Text()
		numbers := strings.Split(text, " ")

		log.Printf("Numbers: %v", text)

		previousVal := -1
		prePreviousVal := -1
		direction := 0
		ignoreCount := 0
		safe := true
		for _, val := range numbers {
			valNum, err := strconv.Atoi(val)

			if err != nil {
				log.Fatal(err)
			}

			if previousVal == -1 {
				previousVal = valNum
				continue
			} else if direction == 0 {
				if previousVal < valNum {
					direction = 1
				} else if previousVal > valNum {
					direction = -1
				} else {
					log.Print("No direction determined, break")
					if ignoreCount > 0 {
						safe = false
						break
					}
					ignoreCount += 1
				}
				log.Printf("D: %v", direction)
			}

			diff := valNum - previousVal
			dirDiff := diff * direction
			//log.Printf("Diff: %v", dirDiff)
			if diff == 0 {
				log.Printf("Diff No change: %v", dirDiff)
				if ignoreCount > 0 {
					safe = false
					break
				}
				ignoreCount += 1
			} else if dirDiff < 0 {
				log.Printf("Diff Unsafe Wrong direction: %v", dirDiff)
				if ignoreCount > 0 {
					safe = false
					break
				}
				ignoreCount += 1
			} else if dirDiff > 3 {
				log.Printf("Diff Unsafe: %v", dirDiff)
				if ignoreCount > 0 {
					safe = false
					break
				}
				ignoreCount += 1
			} else {
				//only change the previous val if no failure
				previousVal = valNum
			}
		}

		if safe {
			log.Print("Safe run detected.")
			successCount += 1
		} else {
			log.Print("Unsafe run detected!")
		}
	}
	log.Printf("Safe runs detected: %v", successCount)
}
