package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	dir, err := os.Getwd()
	handleError(err)

	r, err := git.PlainOpen(dir)
	handleError(err)

	w, err := r.Worktree()
	handleError(err)

	status, err := w.Status()
	handleError(err)

	for path, _ := range status {
		fmt.Println("Please enter a commit message for the file:", path)

		key, err := reader.ReadString('\n')
		handleError(err)

		if key != "\n" {
			w.Add(path)
			w.Commit(key, &git.CommitOptions{})
		}
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
