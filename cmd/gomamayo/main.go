package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/yulog/go-gomamayo"

	"github.com/urfave/cli/v3"
)

const name = "gomamayo"

const version = "0.0.4"

var revision = "HEAD"

func doAnalyze(ctx context.Context, cmd *cli.Command) error {
	d, err := selectDict(cmd.String("sysdict"))
	if err != nil {
		return err
	}
	r := gomamayo.New(d, !cmd.Bool("disable-ignore")).Analyze(cmd.Args().First())
	// fmt.Printf("%+v\n", r)
	obj, err := json.Marshal(r)
	if err != nil {
		return err
	}
	fmt.Println(string(obj))
	return nil
}

func doAddIgnore(ctx context.Context, cmd *cli.Command) error {
	err := gomamayo.AddIgnoreWord(cmd.Args().First())
	if err != nil {
		return err
	}
	fmt.Println("Add ignore word:", cmd.Args().First())

	return nil
}

func doRemoveIgnore(ctx context.Context, cmd *cli.Command) error {
	if cmd.Bool("all") {
		err := gomamayo.RemoveAllIgnoreWords()
		if err != nil {
			return err
		}
		fmt.Println("Remove all ignore word")
		return nil
	}

	err := gomamayo.RemoveIgnoreWord(cmd.Args().First())
	if err != nil {
		return err
	}
	fmt.Println("Remove ignore word:", cmd.Args().First())

	return nil
}

func doListIgnore(ctx context.Context, cmd *cli.Command) error {
	err := gomamayo.ListIgnoreWords(os.Stdout)
	if err != nil {
		return err
	}

	return nil
}

func doImportIgnore(ctx context.Context, cmd *cli.Command) error {
	err := gomamayo.ImportIgnoreWords(cmd.Args().First())
	if err != nil {
		return err
	}
	fmt.Println("Import ignore word")

	return nil
}

func doExportIgnore(ctx context.Context, cmd *cli.Command) error {
	err := gomamayo.ExportIgnoreWords(cmd.Args().First())
	if err != nil {
		return err
	}
	fmt.Println("Export ignore word")

	return nil
}

func main() {
	cmd := &cli.Command{
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
						Name:  "disable-ignore",
						Usage: "disable ignore word",
					},
					&cli.StringFlag{
						Name:  "sysdict",
						Usage: "select dict(ipa,neo,uni,uni3)",
						Value: "neo",
					},
				},
				Action: doAnalyze,
			},
			{
				Name:    "ignoreWord",
				Aliases: []string{"ignore"},
				Usage:   "operations for ignore word",
				Commands: []*cli.Command{
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
								Usage: "remove all ignore word",
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

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
