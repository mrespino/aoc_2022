package main

import (
	"bufio"
	"os"
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

}
