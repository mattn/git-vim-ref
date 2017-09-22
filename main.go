package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/skratchdot/open-golang/open"
)

type Tag struct {
	Ref    string `json:"ref"`
	URL    string `json:"url"`
	Object struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"object"`
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	resp, err := http.Get("https://api.github.com/repos/vim/vim/git/refs/tags")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var tags []Tag
	err = json.NewDecoder(resp.Body).Decode(&tags)
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range tags {
		for _, arg := range os.Args[1:] {
			if tag.Ref == "refs/tags/"+arg {
				url := "https://github.com/vim/vim/commit/" + tag.Object.Sha
				open.Start(url)
			}
		}
	}
}
