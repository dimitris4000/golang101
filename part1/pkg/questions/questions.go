package questions

import (
	"fmt"
	"log"
	"net/url"
)

// Question struct will hold our question data in one entity
type Question struct {
	Question, CorrectAnswer string
}

func decodeURL(encodedValue string) string {
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprint(decodedValue)
}

// GetQuestions will return a list with all the questions to be asked
func GetQuestions() []Question {
	result := []Question{
		{decodeURL("Linus%20Torvalds%20created%20Linux%20and%20Git."), "true"},
		{decodeURL("The%20programming%20language%20%22Python%22%20is%20based%20off%20a%20modified%20version%20of%20%22JavaScript%22."), "false"},
		{decodeURL("The%20logo%20for%20Snapchat%20is%20a%20Bell."), "false"},
		{decodeURL("Ada%20Lovelace%20is%20often%20considered%20the%20first%20computer%20programmer."), "true"},
		{decodeURL("%22HTML%22%20stands%20for%20Hypertext%20Markup%20Language."), "true"},
		{decodeURL("Time%20on%20Computers%20is%20measured%20via%20the%20EPOX%20System."), "false"},
		{decodeURL("The%20Windows%207%20operating%20system%20has%20six%20main%20editions."), "true"},
		{decodeURL("The%20Windows%20ME%20operating%20system%20was%20released%20in%20the%20year%202000."), "true"},
		{decodeURL("The%20NVidia%20GTX%201080%20gets%20its%20name%20because%20it%20can%20only%20render%20at%20a%201920x1080%20screen%20resolution."), "false"},
		{decodeURL("Linux%20was%20first%20created%20as%20an%20alternative%20to%20Windows%20XP."), "false"},
	}
	return result
}
