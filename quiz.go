package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func quiz(timeout <-chan time.Time) (int, int) {
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
		answer, err := strconv.Atoi(row[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(equation)
		fmt.Print("Enter answer: ")
		answerCh := make(chan int, 1)
		go func() {
			var in int
			fmt.Scan(&in)
			answerCh <- in
		}()
		select {
		case <-timeout:
			return numCorrect, (numQuestions + 1)
		case input := <-answerCh:
			if input == answer {
				numCorrect++
				numQuestions++
			}
		}
	}
	return numCorrect, numQuestions
}

func main() {

	start := time.Now()
	timeout := time.After(3 * time.Second)

	// numCorrect, numQuestions := go quiz()
	var numCorrect, numQuestions int
	numCorrect, numQuestions = quiz(timeout)
	end := time.Since(start)
	fmt.Printf("You got %d / %d correct\n", numCorrect, numQuestions)
	fmt.Printf("and took %v seconds\n", end)
}
