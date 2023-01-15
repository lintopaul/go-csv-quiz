package main

import (
	"fmt"
	"time"
)

// Quiz structure
type Quiz struct {
	Question string
	Answer   string
}

// ProcessQuiz processes the quiz
func ProcessQuiz(lines [][]string) int {
	correctAnswers := 0
	timer := time.NewTimer(30 * time.Second)

	for _, line := range lines {
		data := Quiz{
			Question: line[0],
			Answer:   line[1],
		}

		fmt.Println("Question:", data.Question)

		userAnswer := ""
		answered := false
		for !answered {
			select {
			case <-timer.C:
				fmt.Println("Time's up!")
				answered = true
			default:
				fmt.Print("Answer: ")
				_, err := fmt.Scan(&userAnswer)
				if err != nil {
					fmt.Println("Error reading input:", err)
				} else {
					answered = true
				}
			}
		}

		if userAnswer == data.Answer {
			correctAnswers++
		}
		timer.Reset(30 * time.Second)
	}

	return correctAnswers
}
