package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	question, answer string
}

func main() {

	filename := flag.String("csv", "problems.csv", "CSV file")

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Problem opening file, %s", err.Error())
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalf("Problem opening file, %s", err.Error())
	}

	var problems = make([]problem, 0)
	for _, line := range lines {
		problems = append(problems, problem{question: line[0], answer: line[1]})
	}


	score := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, problem.question)

		var userAnswer string
		fmt.Scan(&userAnswer)

		if userAnswer == problem.answer {
			score++
		}
	}

	fmt.Printf("Total correct: %d out of %d\n", score, len(problems))
}

