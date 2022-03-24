package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var arrayOfQuestions []Question
var currentQuestionNumber int = 0

type Question struct {
	name   string
	answer string
}

func display_question() {
	var loadedQandA = arrayOfQuestions[currentQuestionNumber]
	var currentQuestion = loadedQandA.name
	var currentAnswer = loadedQandA.answer
	fmt.Println(currentQuestion + " " + currentAnswer)
}

func load_questions() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data); i++ {
		d0 := data[i]
		question := d0[0]
		answer := d0[1]
		
		var newquestion Question
		newquestion.name = question
		newquestion.answer = answer
		arrayOfQuestions = append(arrayOfQuestions, newquestion)
	}
}

func main() {
	load_questions()
	fmt.Println(arrayOfQuestions)
	display_question()
}
