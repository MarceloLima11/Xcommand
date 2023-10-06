package cmd

import (
	"fmt"
	"os"

	"github.com/MarceloLima11/Xcommand/utils"
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
	utils.ThrowErrorIfHas(err, utils.PathNotFound, 1)

	fileName := fileInfo.Name()
	result := fmt.Sprintln("     ❌ File", fileName, "removed")

	err = os.Remove(filePath)
	utils.ThrowErrorIfHas(err, utils.RemoveFile, 1)

	fmt.Println()
	fmt.Println()
	fmt.Print(result)
	fmt.Println()
	fmt.Println()
	return nil
}
