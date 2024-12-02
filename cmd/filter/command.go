package filter

import (
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/cmd/filter/try"
	"github.com/Kuniwak/name/cmd/filter/validate"
)

var SubCommand = cli.SubCommand{
	Help: "name filter related commands",
	Command: cli.CommandWithSubCommands(
		"filter",
		map[string]cli.SubCommand{
			"try":      try.SubCommand,
			"validate": validate.SubCommand,
		},
	),
}
