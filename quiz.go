package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func quiz() (int, int) {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = 2
	equations, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	numCorrect := 0
	numQuestions := 0
	for _, row := range equations {
		equation := row[0]
		numQuestions += 1
		answer, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(equation)
		var input int
		fmt.Print("Enter answer: ")
		_, err1 := fmt.Scan(&input)
		if err1 != nil {
			log.Fatal(err)
		}
		if answer == input {
			numCorrect += 1
		}
	}
	return numCorrect, numQuestions
}

func main() {
	start := time.Now()
	numCorrect, numQuestions := quiz()
	end := time.Since(start)
	fmt.Printf("You got %d / %d correct\n", numCorrect, numQuestions)
	fmt.Printf("and took %v seconds", end)
}
