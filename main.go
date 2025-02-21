package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	path := addArgument(
		"-csv",
		"a csv file in the format of 'questions,answer' (default: problem.csv)",
		"problem.csv",
	)

	timerLimit := addArgument("-limit", "the time limit for the quizz in seconds (default: 30)", 30)

	showHelp := addArgument("help", "show this help", false)

	parseArgs()

	if *showHelp {
		showDescriptions()
		os.Exit(0)
	}

	columns := readCsvFile(*path)
	quizcards := csvToCards(columns)

	go func() {
		time.Sleep(time.Duration(*timerLimit) * time.Second)
		fmt.Println("\nTime is over!")
		os.Exit(0)
	}()

	takeQuiz(quizcards)
}

type quizCard struct {
	question string
	answer   string
}

func takeQuiz(cards []quizCard) {
	reader := bufio.NewReader(os.Stdin)

	for i, card := range cards {
		fmt.Printf("%v. %v: ", i, card.question)

		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.Trim(userAnswer, "\n ")

		if userAnswer == card.answer {
			fmt.Printf("that's correct!")
		} else {
			fmt.Printf("what the fuck are you doing?")
		}
	}
}

type CSV = [][]string

func csvToCards(v CSV) []quizCard {
	vLen := len(v)

	res := make([]quizCard, vLen)

	for i := 0; i < vLen; i++ {
		res[i] = quizCard{v[i][0], v[i][1]}
	}

	return res
}

func readCsvFile(filePath string) CSV {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, " ", err)
	}

	return records
}
