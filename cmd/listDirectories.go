package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

var (
	ListDirectories = &cli.Command{
		Name:    "allfiles",
		Aliases: []string{"af"},
		Usage:   "List all file",
		Action:  listDirectoriesHandler,
	}

	path string
)

func listDirectoriesHandler(c *cli.Context) error {
	path = c.Args().First()

	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return err
		}
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("     🎃", path)
	fmt.Println()
	fmt.Println()

	files = sortFiles(files) // goroutine (???)
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("📦", file.Name())
		} else {
			fmt.Println("📄", file.Name())
		}
	}
	fmt.Println()
	fmt.Println()

	return nil
}

func sortFiles(files []fs.DirEntry) []fs.DirEntry {
	sort.Slice(files, func(i, j int) bool {
		isDirI := files[i].IsDir()
		isDirJ := files[j].IsDir()

		if isDirI && !isDirJ {
			return true
		} else if !isDirI && isDirJ {
			return false
		}

		return files[i].Name() < files[j].Name()
	})

	return files
}
