package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var arrayOfQuestions []Question
var currentQuestionNumber int = 0
var amountCorrect int = 0

type Question struct {
	name   string
	answer string
}

func display_question() {
	var loadedQandA = arrayOfQuestions[currentQuestionNumber]
	var currentQuestion = loadedQandA.name
	var currentAnswer = loadedQandA.answer
	var usersAnswer string
 	fmt.Println(currentQuestion + " " + currentAnswer)
	fmt.Scanln(&usersAnswer)
	if usersAnswer == currentAnswer {
		amountCorrect++
	}
	currentQuestionNumber++
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
	for{
		display_question()
		if currentQuestionNumber >= len(arrayOfQuestions){
			break
		}
	}
	fmt.Println("You got " + strconv.Itoa(amountCorrect) + " correct out of " + strconv.Itoa(len(arrayOfQuestions)))
}
