package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() error {
	app := &cli.App{
		Name:  "",
		Usage: ">NOME PROVISORIO ⚔<",
		Commands: []*cli.Command{
			ListDirectories,
			ShowDirectoryDetails,
			DeleteFile,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return err
	}

	return nil
}
