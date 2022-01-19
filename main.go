package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:    "CLI",
				Aliases: []string{"c"},
				Usage:   "run on cli",
				Action: func(c *cli.Context) error {
					CLI_main()
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			CLI_main()
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
