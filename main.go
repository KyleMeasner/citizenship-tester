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

	numberOfQuestions := GetNumberToAsk(len(questions))
	AskQuestions(questions, numberOfQuestions)
	termenv.DefaultOutput().ClearScreen()
}

func GetNumberToAsk(numQuestions int) int {
	output := termenv.DefaultOutput()
	output.ClearScreen()

	fmt.Println()
	fmt.Printf("There are %d questions. ", numQuestions)

	number := 0
	for number < 1 || number > 100 {
		fmt.Println("How many questions should I ask you?")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err == nil {
			_, err = fmt.Sscanf(scanner.Text(), "%d", &number)
		}

		if err != nil || number < 1 || number > 100 {
			output.ClearLines(2)
			output.DeleteLines(2)
			fmt.Printf("Please input a number from 1 to %d. ", numQuestions)
		}
	}

	return number
}

func AskQuestions(questions []Question, numberToAsk int) {
	output := termenv.DefaultOutput()

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

		WaitForEnter()

		if len(question.Answers) == 1 {
			fmt.Println(output.String("Answer:").Bold(), question.Answers[0])
		} else {
			fmt.Println(output.String("Possible Answers:").Bold())
			for _, answer := range question.Answers {
				fmt.Println("-", answer)
			}
		}

		WaitForEnter()
		fmt.Println()
	}
}

func WaitForEnter() {
	output := termenv.DefaultOutput()
	fmt.Println()
	fmt.Print(output.String("\n[Enter]").Faint())
	fmt.Scanln()
	output.ClearLines(1)
	output.DeleteLines(1)
}
