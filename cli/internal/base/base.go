package base

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Run       func(cmd *Command, args []string) error
	UsageLine string
	Short     string
	Long      string
	Flag      flag.FlagSet
	Commands  []*Command
}

var CLI = &Command{
	UsageLine: "cli",
	Long:      "CLI is a minimal CLI application with commands",
}

func (c *Command) Lookup(name string) *Command {
	for _, sub := range c.Commands {
		if sub.Name() == name {
			return sub
		}
	}
	return nil
}

// LongName returns the command's long name: all the words in the usage line
// between "cli" and a flag or argument.
func (c *Command) LongName() string {
	name := c.UsageLine
	if i := strings.Index(name, " ["); i >= 0 {
		name = name[:i]
	}
	if name == "cli" {
		return ""
	}
	return strings.TrimPrefix(name, "cli ")
}

func (c *Command) Name() string {
	name := c.LongName()
	if i := strings.LastIndex(name, " "); i >= 0 {
		name = name[i+1:]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "Run 'cli help %s' for details.\n", c.LongName())
	os.Exit(2)
}

var Usage func()
