package cmd

import (
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/cmd/filter"
	"github.com/Kuniwak/name/cmd/info"
	"github.com/Kuniwak/name/cmd/search"
)

var Main = cli.CommandWithSubCommands(
	"name",
	map[string]cli.SubCommand{
		"search": search.SubCommand,
		"info":   info.SubCommand,
		"filter": filter.SubCommand,
	},
)
