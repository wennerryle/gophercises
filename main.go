package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	port := flag.Int("port", 8080, "port of webserver")
	filePath := flag.String("path", "./story.json", "story written in json")
	useCLI := flag.Bool("cli", false, "run in cli mode")

	flag.Parse()

	storyArc, err := parseFromFile(*filePath)

	if err != nil {
		fmt.Printf("An error occured during try to read file: %v\n", err.Error())

		fmt.Printf("Write '%v -h' to see help\n", getExecutableName())
		return
	}

	if *useCLI {
		cliStrategy(storyArc)
		return
	}

	httpStrategy(*port, storyArc)
}

func getExecutableName() string {
	path := os.Args[0]
	i := len(path) - 1

	for ; i != 0; i-- {
		if path[i] == '/' {
			break
		}
	}

	return path[i+1:]
}

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

func httpStrategy(port int, handler http.Handler) {
	url := fmt.Sprintf("localhost:%v", port)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	fmt.Println("server listenning on " + url)
	http.ListenAndServe(url, mux)
}

func (story StoryArc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.URL.Path == "/" {
		storyStarterHTTP(w, r, story)
	}
}

func storyStarterHTTP(w http.ResponseWriter, r *http.Request, story StoryArc) {

}
