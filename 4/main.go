package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	textGrid := [][]string{}

	for scanner.Scan() {
		text := scanner.Text()
		slice := strings.Split(text, "")
		textGrid = append(textGrid, slice)
		//completeText += text
	}

	height, width := len(textGrid), len(textGrid[0])
	word := "MAS"
	//wordLength := len(word)
	found := 0

	/*	for row, rval := range textGrid {
		for col, cval := range rval {
			log.Printf("row: %v col: %v", row, col)
			tests := []string{}
			if col >= wordLength-1 { //can search left letters
				if row >= wordLength-1 { //can search left/up letters
					test := cval
					test += textGrid[row-1][col-1]
					test += textGrid[row-2][col-2]
					test += textGrid[row-3][col-3]
					tests = append(tests, test)
				}
				if row <= height-(wordLength) { // can search left/down letters
					test := cval
					test += textGrid[row+1][col-1]
					test += textGrid[row+2][col-2]
					test += textGrid[row+3][col-3]
					tests = append(tests, test)
				}

				//search left letters
				test := cval
				test += textGrid[row][col-1]
				test += textGrid[row][col-2]
				test += textGrid[row][col-3]
				tests = append(tests, test)
			}
			if col <= width-(wordLength) { //can search right letters
				if row >= wordLength-1 { //can search right/up letters
					test := cval
					test += textGrid[row-1][col+1]
					test += textGrid[row-2][col+2]
					test += textGrid[row-3][col+3]
					tests = append(tests, test)
				}
				if row <= height-(wordLength) { // can search right/down letters
					test := cval
					test += textGrid[row+1][col+1]
					test += textGrid[row+2][col+2]
					test += textGrid[row+3][col+3]
					tests = append(tests, test)
				}
				//search right letters
				test := cval
				test += textGrid[row][col+1]
				test += textGrid[row][col+2]
				test += textGrid[row][col+3]
				tests = append(tests, test)

			}
			if row >= wordLength-1 { //can search up
				test := cval
				test += textGrid[row-1][col]
				test += textGrid[row-2][col]
				test += textGrid[row-3][col]
				tests = append(tests, test)
			}
			if row <= height-(wordLength) { //can search down
				test := cval
				test += textGrid[row+1][col]
				test += textGrid[row+2][col]
				test += textGrid[row+3][col]
				tests = append(tests, test)
			}

			found += searchWord(tests, word)
			log.Printf("found so far: %v", found)
		}
	}*/

	for row, rval := range textGrid {
		if row == 0 || row == height-1 {
			continue
		}
		for col, cval := range rval {
			if col == 0 || col == width-1 {
				continue
			}
			log.Printf("row: %v col: %v", row, col)
			tests := []string{}
			tests2 := []string{}

			if cval == "A" {
				test := textGrid[row-1][col-1]
				test += cval
				test += textGrid[row+1][col+1]
				tests = append(tests, test)
				test = textGrid[row+1][col+1]
				test += cval
				test += textGrid[row-1][col-1]
				tests = append(tests, test)
				if searchWord(tests, word) > 0 {
					test := textGrid[row-1][col+1]
					test += cval
					test += textGrid[row+1][col-1]
					tests2 = append(tests2, test)
					test = textGrid[row+1][col-1]
					test += cval
					test += textGrid[row-1][col+1]
					tests2 = append(tests2, test)
					if searchWord(tests2, word) > 0 {
						found += 1
					}
				}
			}

			log.Printf("found so far: %v", found)
		}
	}

	//hardy har har, throw it all out and try again for part B

	fmt.Printf("Found: %v", found)
	//2657
	//2662

}

func searchWord(list []string, word string) int {
	ret := 0
	for _, val := range list {
		if word == val {
			ret++
		}
	}
	return ret
}
