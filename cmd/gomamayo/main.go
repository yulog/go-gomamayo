package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yulog/go-gomamayo"

	"github.com/urfave/cli/v2"
)

func doAnalyze(cCtx *cli.Context) error {
	r := gomamayo.Analyze(cCtx.Args().First())
	// fmt.Printf("%+v\n", r)
	obj, err := json.Marshal(r)
	if err != nil {
		return err
	}
	fmt.Println(string(obj))
	return nil
}

func main() {
	app := &cli.App{
		Name:        "gomamayo",
		Usage:       "gomamayo analyzer",
		Description: "gomamayo analyzer",
		Commands: []*cli.Command{
			{
				Name:    "analyze",
				Aliases: []string{"a"},
				Usage:   "analyze input string",
				Action:  doAnalyze,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
