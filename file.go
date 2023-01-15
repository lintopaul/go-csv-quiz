package main

import (
	"encoding/csv"
	"os"
)

// ProcessFile parses CSV file
func ProcessFile(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
