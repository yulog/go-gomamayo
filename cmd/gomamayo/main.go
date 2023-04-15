package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yulog/go-gomamayo"

	"github.com/urfave/cli/v2"
)

func doAnalyze(cCtx *cli.Context) error {
	r := gomamayo.New(cCtx.Bool("ignore")).Analyze(cCtx.Args().First())
	// fmt.Printf("%+v\n", r)
	obj, err := json.Marshal(r)
	if err != nil {
		return err
	}
	fmt.Println(string(obj))
	return nil
}

func doAddIgnore(cCtx *cli.Context) error {
	err := gomamayo.AddIgnoreWord(cCtx.Args().First())
	if err != nil {
		return err
	}

	return nil
}

func doRemoveIgnore(cCtx *cli.Context) error {
	err := gomamayo.RemoveIgnoreWord(cCtx.Args().First())
	if err != nil {
		return err
	}

	return nil
}

func doListIgnore(cCtx *cli.Context) error {
	err := gomamayo.ListIgnoreWord(os.Stdout)
	if err != nil {
		return err
	}

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
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "ignore",
						Usage:       "ignore word",
						Value:       true,
						DefaultText: "true"},
				},
				Action: doAnalyze,
			},
			{
				Name:    "addIgnore",
				Aliases: []string{"add"},
				Usage:   "add ignore word",
				Action:  doAddIgnore,
			},
			{
				Name:    "removeIgnore",
				Aliases: []string{"remove"},
				Usage:   "remove ignore word",
				Action:  doRemoveIgnore,
			},
			{
				Name:    "listIgnore",
				Aliases: []string{"list"},
				Usage:   "list ignore word",
				Action:  doListIgnore,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
