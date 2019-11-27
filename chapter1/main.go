package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func decodeURL(encodedValue string) string {
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(decodedValue)
}

func getQuestions() ([]string, []bool) {
	correctAnswers := []bool{
		true,
		false,
		false,
		true,
		true,
		false,
		true,
		true,
		false,
		false,
	}
	questions := []string{
		decodeURL("Linus%20Torvalds%20created%20Linux%20and%20Git."),
		decodeURL("The%20programming%20language%20%22Python%22%20is%20based%20off%20a%20modified%20version%20of%20%22JavaScript%22."),
		decodeURL("The%20logo%20for%20Snapchat%20is%20a%20Bell."),
		decodeURL("Ada%20Lovelace%20is%20often%20considered%20the%20first%20computer%20programmer."),
		decodeURL("%22HTML%22%20stands%20for%20Hypertext%20Markup%20Language."),
		decodeURL("Time%20on%20Computers%20is%20measured%20via%20the%20EPOX%20System."),
		decodeURL("The%20Windows%207%20operating%20system%20has%20six%20main%20editions."),
		decodeURL("The%20Windows%20ME%20operating%20system%20was%20released%20in%20the%20year%202000."),
		decodeURL("The%20NVidia%20GTX%201080%20gets%20its%20name%20because%20it%20can%20only%20render%20at%20a%201920x1080%20screen%20resolution."),
		decodeURL("Linux%20was%20first%20created%20as%20an%20alternative%20to%20Windows%20XP."),
	}
	return questions, correctAnswers
}

func main() {
	questions, correctAnswers := getQuestions()

	successAnswerCount := 0
	for i, q := range questions {
		fmt.Printf("Question #%d: %s=", i+1, q)

		answer := ""
		fmt.Scanf("%s\n", &answer)
		if strings.ToLower(strings.Trim(answer, "\n ")) == fmt.Sprintf("%t", correctAnswers[i]) {
			successAnswerCount++
		}
	}
	println("You scored", successAnswerCount, "out of", len(questions))

}
