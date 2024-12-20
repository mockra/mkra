package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	r, _ := git.PlainOpen(".")
	w, _ := r.Worktree()

	status, _ := w.Status()

	for path, _ := range status {
		fmt.Println("Please enter a commit message for the file:", path)
		key, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if key != "\n" {
			w.Add(path)
			w.Commit(key, &git.CommitOptions{})
		}
	}
}
