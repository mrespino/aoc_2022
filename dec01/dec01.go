package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// errorHandler is a quick and dirty way to check err after function calls
func errorHandler(e error) error {
	if e != nil {
		return e
	}

	return nil
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
	for i, line := range data {
		if line != "" {
			foodCal, cerr := strconv.Atoi(line)

			errorHandler(cerr)

			totalCalories += foodCal

			log.Println("elf: " + strconv.Itoa(elf) + ", meal: " + strconv.Itoa(i) + ", calories: " + line + "\n")

		} else {
			elfCalories = append(elfCalories, totalCalories)
			log.Println("Total for elf: " + strconv.Itoa(elf) + "calories carried:" + strconv.Itoa(elfCalories[elf]))
			totalCalories = 0
			elf++
		}
	}

	// now find which elf is carrying the most calories
	for j := 1; j < elf; j++ {
		if elfCalories[0] < elfCalories[j] {
			elfCalories[0] = elfCalories[j]
		}
	}

	log.Println("the most calories carried by an elf is: ", elfCalories[0])
}
