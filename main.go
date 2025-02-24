package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
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
		http.Redirect(w, r, "/intro", http.StatusTemporaryRedirect)
		return
	}

	storyStarterHTTP(w, r, story, getTemplate())
}

func storyStarterHTTP(w http.ResponseWriter, r *http.Request, story StoryArc, template *template.Template) {
	path := r.URL.Path[1:]
	storyPart, ok := story[path]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	template.Execute(w, storyPart)
}

func getTemplate() *template.Template {
	t, _ := template.New("page").Parse(
		`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>My Adventure</title>
  <script src="https://unpkg.com/@tailwindcss/browser@4"></script>
</head>
<body class="max-w-4xl mx-auto px-4">
  <h1 class="bg-blue-200 h-10 mb-4 mx-auto pt-2 rounded-b-lg text-center">
    {{.Title}}
  </h1>
  <div class="flex flex-col gap-y-2">
    {{
      range .Story
    }}
      <p class="indent-12">
        {{.}}
      </p>
    {{end}}
  </div>
  <div class="flex flex-col gap-y-2 mt-4">
    {{
      range .Options
    }}
    <a href="/{{.Arc}}" class="bg-blue-100 hover:bg-blue-200 mt-2 outline-1 outline-gray-400 p-2 rounded-sm transition-colors">
      {{.Text}}
    </a>
    {{end}}
  </div>
</body>
</html>`,
	)

	return t
}
