package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read line by line into memory
// all file contents is stores in lines[]
func readLines(path string) ([]string, error) {
	//check the file path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//close the file
	defer file.Close()
	//var to the string
	var lines []string
	//to scan the file
	scanner := bufio.NewScanner(file)
	// log.Logger(scanner.Text())
	for scanner.Scan() {
		// add the file to the var string --line
		lines = append(lines, scanner.Text())
	}
	//Return the value of the scanner
	return lines, scanner.Err()
}

func main() {
	// open file for reading
	// read line by line

	//call the readLines function to read the file
	lines, err := readLines("read.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	// print file contents
	// For loop generate each line
	for i, line := range lines {
		fmt.Println(i, line)
	}
}
