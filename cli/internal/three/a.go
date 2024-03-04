package three

import (
	"fmt"

	"github.com/jroimartin/template/cli/internal/base"
)

var cmdA = &base.Command{
	UsageLine: "cli three a [-f]",
	Short:     "short description for a",
	Long: `
Long description for three a.

It supports multiple lines.
	`,
}

var exampleFlag bool

func init() {
	cmdA.Run = runA
	cmdA.Flag.BoolVar(&exampleFlag, "f", false, "Example flag")
}

func runA(cmd *base.Command, args []string) error {
	fmt.Printf("three a! exampleFlag=%v\n", exampleFlag)
	return nil
}
