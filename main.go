package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question, answer string
}

func main() {

	filename := flag.String("csv", "problems.csv", "The CSV file")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Problem opening file, %s", err.Error())
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("Problem reading file, %s", err.Error())
	}

	var problems = make([]problem, 0)
	for _, line := range lines {
		problems = append(problems, problem{question: line[0], answer: line[1]})
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerChan := make(chan string)

	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, problem.question)

		go func() {
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerChan <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime up!")
			fmt.Printf("Total correct: %d out of %d\n", score, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.answer {
				score++
			}
		}
	}

	fmt.Printf("Total correct: %d out of %d\n", score, len(problems))
}
