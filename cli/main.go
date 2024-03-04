// cli is a reference for CLI applications based on the Go tool.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jroimartin/template/cli/internal/base"
	"github.com/jroimartin/template/cli/internal/help"
	"github.com/jroimartin/template/cli/internal/one"
	"github.com/jroimartin/template/cli/internal/three"
	"github.com/jroimartin/template/cli/internal/two"
)

func init() {
	base.CLI.Commands = []*base.Command{
		one.CmdOne,
		two.CmdTwo,
		three.CmdThree,
	}

	base.Usage = mainUsage
}

func main() {
	log.SetFlags(0)

	flag.Usage = base.Usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		base.Usage()
	}

	if args[0] == "help" {
		help.Help(args[1:])
		return
	}

	cmd, used := lookupCmd(args)
	if len(cmd.Commands) == 0 {
		invoke(cmd, args[used-1:])
	} else {
		if used >= len(args) {
			//Subcommand is missing.
			help.PrintUsage(cmd)
			os.Exit(2)
		}

		// Command or subcommand are unknown.
		helpArg := ""
		if used > 0 {
			// Subcommand is unknown.
			helpArg += " " + strings.Join(args[:used], " ")
		}

		fmt.Fprintf(os.Stderr, "cli %s: unknown command\nRun 'cli help%s' for usage.\n", args[0], helpArg)
		os.Exit(2)
	}
}

func lookupCmd(args []string) (cmd *base.Command, used int) {
	cmd = base.CLI
	for used < len(args) {
		c := cmd.Lookup(args[used])
		if c == nil {
			// Command not found. If len(cmd.Commands) == 0 it may be an
			// argument for the current command.
			break
		}
		cmd = c
		used++
	}
	return cmd, used
}

func invoke(cmd *base.Command, args []string) {
	cmd.Flag.Usage = func() { cmd.Usage() }
	cmd.Flag.Parse(args[1:])
	args = cmd.Flag.Args()
	if err := cmd.Run(cmd, args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func mainUsage() {
	help.PrintUsage(base.CLI)
	os.Exit(2)
}
