package main

import (
	"log"
	"net/http"
)

func QuestionFunc(w http.ResponseWriter, r *http.Request) {
	q := NewQuestion(
		"Столица Южной Кореи?",
		[]string{
			"Сеул",
			"Пусан",
			"Пхеньян",
			"Тэгу",
		},
		0)
	data, err := ConvertQuestionToJson(q)
	if err != nil {
		log.Printf("Error with marshal of a question: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}
	w.Write(data)
}

func QuestionsListFunc(w http.ResponseWriter, r *http.Request) {
	q := []string{"ОТУ", "ЦОС", "История", "Философия"}
	data, err := ConvertQuestionsListToJson(q)
	if err != nil {
		log.Printf("Error with marshal of a question: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal Server Error"))
		return
	}
	w.Write(data)
}
