package main

import (
	"fmt"
	"strconv"
	"strings"
	"thequizgame/pkg/questions"
)

func main() {
	questions := questions.GetQuestions()
	successAnswerCount := 0

	// Questions number prompt
	fmt.Printf("How many questions would you prefer?\n")
	answer := ""
	fmt.Scanf("%s\n", &answer)
	questionsToAnswer, err := strconv.Atoi(answer)
	if err != nil || questionsToAnswer > len(questions) {
		fmt.Printf("Invalid input. Using the default(i.e. %d)\n", len(questions))
		questionsToAnswer = len(questions)
	}

	// Ask the questions
	for i, q := range questions {
		if i >= questionsToAnswer {
			break
		}

		fmt.Printf("Question #%d: %s=", i+1, q.Question)

		answer := ""
		fmt.Scanf("%s\n", &answer)
		if strings.ToLower(strings.Trim(answer, "\n ")) == q.CorrectAnswer {
			successAnswerCount++
		}
	}

	// Print summary
	println("You scored", successAnswerCount, "out of", questionsToAnswer)
}
