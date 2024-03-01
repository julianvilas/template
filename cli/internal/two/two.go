package two

import (
	"fmt"

	"github.com/jroimartin/template/cli/internal/base"
)

var CmdTwo = &base.Command{
	UsageLine: "two [-f]",
	Short:     "short description for two",
	Long: `
Long description for two.

It supports multiple lines.
	`,
}

var exampleFlag bool

func init() {
	CmdTwo.Run = runTwo
	CmdTwo.Flag.BoolVar(&exampleFlag, "f", false, "Example flag")
}

func runTwo(cmd *base.Command, args []string) error {
	fmt.Printf("two! exampleFlag=%v\n", exampleFlag)
	return nil
}
