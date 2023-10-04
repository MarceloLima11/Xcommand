package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	ShowDirectoryDetails = &cli.Command{
		Name:    "filedetails",
		Aliases: []string{"dt"},
		Usage:   "Show details of a specific file",
		Action:  showDirectoryDetails,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "external",
				Value: "",
				Usage: "Specifies an external file to show details for",
			},
		},
	}

	filePath string
)

func showDirectoryDetails(c *cli.Context) error {
	externalFilePath := c.String("external")
	if externalFilePath != "" {
		filePath = externalFilePath
	} else {
		fileName := c.Args().Get(0)
		if fileName == "" {
			return cli.Exit("Missed <Name>", 1)
		}
		currentFolder, err := os.Getwd()
		if err != nil {
			return cli.Exit(err, 1)
		}
		filePath = filepath.Join(currentFolder, fileName)
	}

	fileInfo, err := os.Stat(filePath)
	if fileInfo == nil {
		return cli.Exit(err.Error(), 1)
	}
	fmt.Println()
	fmt.Println("     🕵️  File Details: ")

	fmt.Println()
	fmt.Println()
	fmt.Println("⭐️ ", fileInfo.Name())
	fmt.Println("📏 ", fileInfo.Size())
	fmt.Println("🕒 ", fileInfo.ModTime().Format(time.DateTime))

	return nil
}
