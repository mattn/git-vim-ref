package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

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

func openCommitPage(sha string) {
	url := "https://github.com/vim/vim/commit/" + sha
	open.Start(url)
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	name := os.Args[1]
	b, err := exec.Command("git", "rev-list", "-n", "1", name).CombinedOutput()
	if err == nil && len(b) > 0 {
		openCommitPage(string(b))
		return
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
		if tag.Ref == "refs/tags/"+name {
			openCommitPage(tag.Object.Sha)
		}
	}
}
