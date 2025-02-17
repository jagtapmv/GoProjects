package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fileFlag := flag.String("csvpath", "./assets/problems.csv", "Provide the CSV file path for problems")
	flag.Parse()

	fileContent, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("Error occured while opening csv file")
		return
	}
	actualContent, err := csv.NewReader(fileContent).ReadAll()
	if err != nil {
		fmt.Println("Error occured while reading csv file")
		return
	}

	problems := parseLines(actualContent)

	timer := time.NewTimer(5 * time.Second)
	correct := 0
mainloop:
	for _, i := range problems {
		fmt.Println("What is ", i.p)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break mainloop
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == strings.TrimSpace(i.a) {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(actualContent))
}

type problem struct {
	p string
	a string
}

func parseLines(problems [][]string) []problem {
	result := make([]problem, len(problems))

	for i, line := range problems {
		result[i] = problem{
			p: line[0],
			a: line[1],
		}
	}
	return result
}
