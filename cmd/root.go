package cmd

import (
	"os"

	"github.com/MarceloLima11/Xcommand/utils"
	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "Xcommand",
		Usage: "Terminal directory manager",
		Commands: []*cli.Command{
			ListDirectories,
			ShowDirectoryDetails,
			DeleteFile,
		},
	}

	err := app.Run(os.Args)
	utils.ThrowErrorIfHas(err, utils.BasicError, 1)

	return nil
}
