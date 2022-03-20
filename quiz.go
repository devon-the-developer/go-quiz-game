package main

import "fmt"
import "encoding/csv"
import "os"
import "log"

func display_question(){
	fmt.Println("Here is a question for you")
}

func load_questions(){
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
	
//	questionsList := createQuestionsList(data)

	fmt.Printf("%+v\n", data)
}

func main() {
	//display_question()
	load_questions()
}