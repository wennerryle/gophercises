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

	go func() {
		time.Sleep(time.Duration(*timerLimit) * time.Second)
		fmt.Println("\nTime is over!")
		os.Exit(1)
	}()

	parseArgs()

	if *showHelp {
		showDescriptions()
		os.Exit(1)
	}

	columns := readCsvFile(*path)
	reader := bufio.NewReader(os.Stdin)

	for i, v := range columns {
		question := v[0]
		answer := v[1]

		fmt.Printf("%v. %v: ", i, question)

		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.Trim(userAnswer, "\n ")

		if userAnswer == answer {
			println("that's correct!")
		} else {
			println("what the fuck are you doing?")
		}
	}
}

func readCsvFile(filePath string) [][]string {
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
