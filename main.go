package main

import (
	"os"

	"github.com/golangaccount/cli/command"
)

func main() {
	app.Commands = command.GetCommands()
	app.Run(os.Args)
}
