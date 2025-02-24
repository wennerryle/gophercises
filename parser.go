package main

import (
	"encoding/json"
	"os"
)

type StoryOption struct {
	Text string
	Arc  string
}

type StoryPart struct {
	Title   string
	Story   []string
	Options []StoryOption
}

type StoryArc map[string]StoryPart

func parseFromFile(path string) (StoryArc, error) {
	var story StoryArc

	bytes, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &story)

	return story, err
}
