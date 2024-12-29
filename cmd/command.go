package cmd

import (
	"github.com/Kuniwak/name/cli"
	"github.com/Kuniwak/name/cmd/filter"
	"github.com/Kuniwak/name/cmd/info"
	"github.com/Kuniwak/name/cmd/search"
	"github.com/Kuniwak/name/cmd/validate"
	"github.com/Kuniwak/name/cmd/yomi"
)

var Main = cli.CommandWithSubCommands(
	"name",
	map[string]cli.SubCommand{
		"search":   search.SubCommand,
		"info":     info.SubCommand,
		"filter":   filter.SubCommand,
		"validate": validate.SubCommand,
		"yomi":     yomi.SubCommand,
	},
)
