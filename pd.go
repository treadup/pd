package main

import (
    "fmt"
	"os"
	"path/filepath"
)

// isProjectDir determines if the given directory dir is
// a project directory. Currently a directory is a project
// directory if it contains a .git folder. Returns true if
// the directory is a project directory. Returns false
// otherwise.
func isProjectDir(dir string) bool {
	gitDir := filepath.Join(dir, ".git")
    _, err := os.Stat(gitDir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

// findProjectDir finds the project directory that the given
// directory dir is contained in. Walks from the given dir
// up to the root directory. If a project directory is found
// projectDir will contain the project directory and ok will
// be true. If a project directory is not found projectDir
// will contain the empty string and ok will be false.
func findProjectDir(dir string) (projectDir string, ok bool) {
	for dir != "/" {
		if isProjectDir(dir) {
			return dir, true
		} else {
			dir = filepath.Dir(dir)
		}
	}
	return "", false
}

func main() {
    dir, err := os.Getwd()
	if err != nil {
		os.Exit(128)
	}
	projectDir, ok := findProjectDir(dir)
	if ok {
		fmt.Print(projectDir)
	} else {
		fmt.Fprintln(os.Stderr, "Not in a project.")
		os.Exit(1)
	}
}
