package main

import (
	"time"

	"github.com/urfave/cli"
)

var app *cli.App

func init() {
	app = &cli.App{}
	app.Name = "flypig"
	app.Version = "1.00.00"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "flypig",
			Email: "379028604@qq.com",
		},
	}
	app.Copyright = ""
	app.HelpName = ""
	app.Usage = ""
	app.UsageText = ""
	app.ArgsUsage = ""
}
