package exists

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golangaccount/cli/command"
	"github.com/urfave/cli"
)

func init() {
	command.PushCommand(&cli.Command{
		Name:        "exists",
		ShortName:   "exists",
		Aliases:     []string{"exists"},
		Usage:       "判断是否存在",
		UsageText:   "exist -f ",
		Description: "",
		ArgsUsage:   "-f {file} -dir {dir}",
		Category:    "",
		Action: func(c *cli.Context) error {
			if strings.TrimSpace(c.String("f")) != "" {
				exists(c.String("f"), "f", c.App.Writer)
			} else if strings.TrimSpace(c.String("dir")) != "" {
				exists(c.String("dir"), "dir", c.App.Writer)
			} else if len(c.Args()) > 0 {
				exists(c.Args()[0], "f", c.App.Writer)
			} else {
				fmt.Fprintln(c.App.Writer, "false")
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "f"},
			cli.StringFlag{Name: "dir"},
		},
	})
}

func exists(path string, tp string, w io.Writer) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(w, "false")
		return
	}
	if info.IsDir() {
		if tp == "dir" {
			fmt.Fprintln(w, "true")
		} else {
			fmt.Fprintln(w, "true; path is a file")
		}
	} else {
		if tp == "f" {
			fmt.Fprintln(w, "true")
		} else {
			fmt.Fprintln(w, "true; path is a dir")
		}
	}
}
