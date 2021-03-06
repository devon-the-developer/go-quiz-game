// Initial try at programming a quiz using Go without guide
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var arrayOfQuestions []Problem
var currentQuestionNumber int = 0
var amountCorrect int = 0

type Problem struct {
	question   string
	answer string
}

func display_question() {
	loadedQandA := arrayOfQuestions[currentQuestionNumber]
	currentQuestion := loadedQandA.question
	currentAnswer := loadedQandA.answer
	var usersAnswer string
 	fmt.Println(currentQuestion + " ")
	fmt.Scanln(&usersAnswer)
	if usersAnswer == currentAnswer {
		amountCorrect++
	}
	currentQuestionNumber++
}

func load_questions() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("failed to open the CSV file %s\n", *csvFilename)
		os.Exit(1)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data); i++ {
		d0 := data[i]
		question := d0[0]
		answer := d0[1]
		
		var newProblem Problem
		newProblem.question = question
		newProblem.answer = answer
		arrayOfQuestions = append(arrayOfQuestions, newProblem)
	}
}

func main() {
	load_questions()
	// Create timer and start counting
	timer := time.NewTimer(30 * time.Second)
	// Because the breaks are nested we need to create the label "Loop" to break out of
	// Problem with this loop is it doesn't break before a user finishes answering -> need someway to check during display_question()
	Loop:
		for{
			select{
			case <-timer.C:
				println("Time is up: 30 seconds")
				break Loop
			default:
				if currentQuestionNumber >= len(arrayOfQuestions){
					break Loop
				}
				display_question()
			}
		}
	// Print out how many correct
	fmt.Println("You got " + strconv.Itoa(amountCorrect) + " correct out of " + strconv.Itoa(len(arrayOfQuestions)))
}
