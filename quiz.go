package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

var arrayOfQuestions []Question

type Question struct {
	name   string
	answer string
}

func display_question() {
	fmt.Println("Here is a question for you")
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
		fmt.Println(newquestion)
		//arrayOfQuestions.append(newquestion)
	}
}

func main() {
	//display_question()
	load_questions()
}
