package three

import (
	"github.com/jroimartin/template/cli/internal/base"
)

var CmdThree = &base.Command{
	UsageLine: "cli three",
	Short:     "three has subcommands",
	Long: `
Long description for three.

It supports multiple lines.
	`,

	Commands: []*base.Command{
		cmdA,
		cmdB,
	},
}
