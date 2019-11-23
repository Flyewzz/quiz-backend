package main

import "encoding/json"

type Question struct {
	Title   string   `json:"title"`
	Answers []string `json:"answers"`
	Right   int      `json:"right"`
}

func ConvertQuestionToJson(q Question) ([]byte, error) {
	result, err := json.Marshal(q)
	return result, err
}

func NewQuestion(title string, answers []string, right int) Question {
	return Question{
		Title:   title,
		Answers: answers,
		Right:   right,
	}
}
