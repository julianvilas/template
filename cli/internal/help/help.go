package help

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/jroimartin/template/cli/internal/base"
)

func Help(args []string) {
	cmd := base.CLI

Args:
	for i, arg := range args {
		for _, sub := range cmd.Commands {
			if sub.Name() == arg {
				cmd = sub
				continue Args
			}
		}

		// helpSuccess is the help command using as many args as possible that
		// would succeed.
		helpSuccess := "cli help"
		if i > 0 {
			helpSuccess += " " + strings.Join(args[:i], " ")
		}
		fmt.Fprintf(os.Stderr, "cli help %s: unknown help topic. Run '%s'.\n", strings.Join(args, " "), helpSuccess)
		os.Exit(2)
	}

	if len(cmd.Commands) > 0 {
		PrintUsage(cmd)
	} else {
		tmpl(helpTemplate, cmd)
	}
}

func PrintUsage(cmd *base.Command) {
	tmpl(usageTemplate, cmd)
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func tmpl(text string, data any) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(os.Stderr, data); err != nil {
		panic(err)
	}
}

const usageTemplate = `{{.Long | trim}}

Usage:

	{{.UsageLine}} <command> [arguments]

The commands are:
{{range .Commands}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "cli help{{with .LongName}} {{.}}{{end}} <command>" for more information about a command.
`

const helpTemplate = `usage: {{.UsageLine}}

{{.Long | trim}}
`
