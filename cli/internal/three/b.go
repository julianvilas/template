package three

import (
	"fmt"

	"github.com/jroimartin/template/cli/internal/base"
)

var cmdB = &base.Command{
	UsageLine: "cli three b [-f]",
	Short:     "short description for b",
	Long: `
Long description for three b.

It supports multiple lines.
	`,
}

func init() {
	cmdB.Run = runB
	cmdB.Flag.BoolVar(&exampleFlag, "f", false, "Example flag")
}

func runB(cmd *base.Command, args []string) error {
	fmt.Printf("three b! exampleFlag=%v\n", exampleFlag)
	return nil
}
