package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

// errorHandler is a quick and dirty way to check err after function calls
func errorHandler(e error) {
	if e != nil {
		panic(e)
	}
}

// fileSlurper reads a file of input data and creates an array of data inputs
func fileSlurper(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)

	errorHandler(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// main is well, the main function that does all of the work.
func main() {
	var elf int = 0

	var totalCalories int = 0

	var elfCalories []int

	data, derr := fileSlurper("input")

	errorHandler(derr)

	// read though the array, totaling each elf's calories
	for _, line := range data {
		if line != "" {
			foodCal, cerr := strconv.Atoi(line)

			errorHandler(cerr)

			totalCalories += foodCal

		} else {
			elfCalories = append(elfCalories, totalCalories)

			totalCalories = 0
			elf++
		}
	}

	sort.Ints(elfCalories)

	lastelf := len(elfCalories)

	log.Println("the most calories carried by an elf is: ", elfCalories[lastelf-1])

	// now find the total calories carried by the top three elves.
	totalThreeCalories := elfCalories[lastelf-1]
	totalThreeCalories += elfCalories[lastelf-2]
	totalThreeCalories += elfCalories[lastelf-3]

	log.Println("the total calories carried by  the top three elves is: ", totalThreeCalories)

}
