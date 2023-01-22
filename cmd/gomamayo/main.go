package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/yulog/go-gomamayo"

	"github.com/urfave/cli/v2"
)

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
				Action: func(cCtx *cli.Context) error {
					r := gomamayo.Analyze(cCtx.Args().First())
					// fmt.Printf("%+v\n", r)
					obj, _ := json.Marshal(r)
					fmt.Println(string(obj))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
