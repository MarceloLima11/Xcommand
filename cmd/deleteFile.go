package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var DeleteFile = &cli.Command{
	Name:    "delete",
	Aliases: []string{"del"},
	Usage:   "Delete a specify file",
	Action:  deleteFile,
}

func deleteFile(c *cli.Context) error {
	filePath := c.Args().First()

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	fileName := fileInfo.Name()
	result := fmt.Sprintln("     ❌ File", fileName, "removed")
	err = os.Remove(filePath)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Print(result)
	fmt.Println()
	fmt.Println()
	return nil
}
