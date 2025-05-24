package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
			fmt.Println("Error:", err)
			return
		}
		if answer == input {
			numCorrect += 1
		}
	}
	fmt.Printf("You got %d / %d correct\n", numCorrect, numQuestions)
}
