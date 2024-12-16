package filter

import (
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/cmd/filter/test"
	"github.com/Kuniwak/name/cmd/filter/validate"
)

var SubCommand = cli.SubCommand{
	Help: "name filter related commands",
	Command: cli.CommandWithSubCommands(
		"filter",
		map[string]cli.SubCommand{
			"test":     test.SubCommand,
			"validate": validate.SubCommand,
		},
	),
}
