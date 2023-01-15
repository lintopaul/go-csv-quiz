package main

import (
	"fmt"
)

// fileName represents the input file for quiz
const (
	fileName = "quiz.csv"
)

func main() {
	lines := ProcessFile(fileName)
	fmt.Println("Number of correct answers:", ProcessQuiz(lines))
}
