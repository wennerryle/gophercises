package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func cliStrategy(storyArc StoryArc) {
	storyPart, ok := storyArc["intro"]

	if !ok {
		fmt.Println("User Intro not defined. Failed to start story")
		return
	}

	for ok {
		path := showConsoleStoryPart(storyPart)
		storyPart, ok = storyArc[path]

		fmt.Println()
	}
}

func showConsoleStoryPart(storyPart StoryPart) string {
	fmt.Println(storyPart.Title)

	fmt.Print(strings.Repeat("_", utf8.RuneCountInString(storyPart.Title)) + "\n\n")

	for _, v := range storyPart.Story {
		fmt.Println(v + "\n")
	}

	optionsAmount := len(storyPart.Options)

	for i := 0; i < optionsAmount; i++ {
		fmt.Printf("%v. %v\n", i+1, storyPart.Options[i].Text)
	}

	if optionsAmount == 0 {
		return ""
	}

	selected := getAnswer(1, len(storyPart.Options))

	return storyPart.Options[selected-1].Arc
}

// Accepts answer in selected range. Ex: getAnswer(0, 1) will return 0 or 1
func getAnswer(min, max int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Твой выбор: ")
		bytes, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println("Произошла ошибка:", err.Error())
			continue
		}

		selected, err := strconv.Atoi(string(bytes))

		if err != nil {
			fmt.Print("Можно вводить только числа. ")
			continue
		}

		if selected > max {
			fmt.Printf("Максимальный вариант ответа %v\n", max)
			continue
		}

		if selected < min {
			fmt.Printf("Минимальный вариант ответа %v\n", min)
			continue
		}

		return selected
	}
}
