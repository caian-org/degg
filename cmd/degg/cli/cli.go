package cli

import (
	// standard
	"context"
	"net/mail"
	"os"

	// 3rd-party
	"github.com/urfave/cli/v3"
)

type (
	CLIActionCallback func(input string, output string) error

	CLI struct {
		cmd *cli.Command
	}
)

const (
	flagInputFile  = "input"
	flagOutputFile = "output"
)

var (
	authors = []any{
		&mail.Address{Name: "Caian Ertl", Address: "hi@caian.org"},
	}

	inputFlag = &cli.StringFlag{
		Name:     flagInputFile,
		Usage:    "the input file containing the enum definition",
		Aliases:  []string{"i"},
		Required: true,
	}

	outputFlag = &cli.StringFlag{
		Name:     flagOutputFile,
		Usage:    "the output file to write the generated code",
		Aliases:  []string{"o"},
		Required: true,
	}
)

func New(callback CLIActionCallback) *CLI {
	cmd := &cli.Command{
		Name:    "degg",
		Usage:   "Dumb Enum Generator for Go",
		Version: programVersion,
		Authors: authors,
		Flags:   []cli.Flag{inputFlag, outputFlag},

		Metadata: map[string]any{
			"compiled": programCompiledAt,
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			return callback(
				cmd.String(flagInputFile),
				cmd.String(flagOutputFile),
			)
		},
	}

	return &CLI{cmd}
}

func (c *CLI) Act() error {
	return c.cmd.Run(context.Background(), os.Args)
}
