package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/muesli/termenv"
)

type Question struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

func main() {
	questionsFile, err := os.Open("questions.json")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	questions := []Question{}

	jsonParser := json.NewDecoder(questionsFile)
	err = jsonParser.Decode(&questions)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Printf("There are %d questions. ", len(questions))

	number := 0
	for number < 1 || number > 100 {
		fmt.Println("How many questions should I ask you?")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err = scanner.Err()
		if err == nil {
			_, err = fmt.Sscanf(scanner.Text(), "%d", &number)
		}
		if err != nil || number < 1 || number > 100 {
			fmt.Printf("\nPlease input a number from 1 to %d. ", len(questions))
		}
	}

	AskQuestions(questions, number)
}

func AskQuestions(questions []Question, numberToAsk int) {
	output := termenv.DefaultOutput()
	enterString := output.String("\n[Enter]").Faint()

	remaining := make([]int, len(questions))
	for i := range remaining {
		remaining[i] = i
	}

	for range numberToAsk {
		output.ClearScreen()

		i := rand.Intn(len(remaining))
		question := questions[remaining[i]]
		remaining = append(remaining[:i], remaining[i+1:]...)

		fmt.Println()
		questionString := output.String("Question:", question.Question).Bold().Italic()
		fmt.Println(questionString)

		fmt.Println()
		fmt.Print(enterString)
		fmt.Scanln()
		output.ClearLines(1)

		if len(question.Answers) == 1 {
			answerString := output.String("Answer:", question.Answers[0]).Bold()
			fmt.Println(answerString)
		} else {
			fmt.Println(output.String("Possible Answers:").Bold())
			for _, answer := range question.Answers {
				fmt.Println("-", answer)
			}
		}

		fmt.Print(enterString)
		fmt.Scanln()
	}
}
