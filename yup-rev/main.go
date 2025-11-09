package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/rev"
)

const (
	flagSeparate = "separate"
)

func main() {
	app := &cli.App{
		Name:  "rev",
		Usage: "reverse lines characterwise",
		UsageText: `rev [OPTIONS] [FILE...]

   The rev utility copies the specified files to standard output, reversing
   the order of characters in every line. If no files are specified, standard
   input is read.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagSeparate,
				Aliases: []string{"s"},
				Usage:   "reverse characters separately for each line",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "rev: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.Bool(flagSeparate) {
		params = append(params, Separate)
	}

	// Create and execute the rev command
	cmd := Rev(params...)
	return gloo.Run(cmd)
}
