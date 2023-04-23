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
	if cCtx.Bool("all") {
		err := gomamayo.RemoveAllIgnoreWords()
		if err != nil {
			return err
		}
		return nil
	}

	err := gomamayo.RemoveIgnoreWord(cCtx.Args().First())
	if err != nil {
		return err
	}

	return nil
}

func doListIgnore(cCtx *cli.Context) error {
	err := gomamayo.ListIgnoreWords(os.Stdout)
	if err != nil {
		return err
	}

	return nil
}

func doImportIgnore(cCtx *cli.Context) error {
	err := gomamayo.ImportIgnoreWords("")
	if err != nil {
		return err
	}

	return nil
}

func doExportIgnore(cCtx *cli.Context) error {
	err := gomamayo.ExportIgnoreWords("")
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
				Name:    "ignoreWord",
				Aliases: []string{"ignore"},
				Usage:   "operations for ignore word",
				Subcommands: []*cli.Command{
					{
						Name:    "add",
						Aliases: []string{"a"},
						Usage:   "add ignore word",
						Action:  doAddIgnore,
					},
					{
						Name:    "remove",
						Aliases: []string{"r"},
						Usage:   "remove ignore word",
						Flags: []cli.Flag{
							&cli.BoolFlag{
								Name:  "all",
								Usage: "remove all",
							},
						},
						Action: doRemoveIgnore,
					},
					{
						Name:    "list",
						Aliases: []string{"l"},
						Usage:   "list ignore word",
						Action:  doListIgnore,
					},
					{
						Name:    "import",
						Aliases: []string{"i"},
						Usage:   "import ignore word",
						Action:  doImportIgnore,
					},
					{
						Name:    "export",
						Aliases: []string{"e"},
						Usage:   "export ignore word",
						Action:  doExportIgnore,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
