package command

import "github.com/chzyer/readline"

var Completer = readline.NewPrefixCompleter(
	// Common command
	readline.PcItem("help"),
	readline.PcItem("exit"),

	// Create
	readline.PcItem("create"),
	readline.PcItem("join"),
	//readline.PcItem("scan"),

	// Game command
	readline.PcItem("say"),
	readline.PcItem("vote"),
	readline.PcItem("kill"),
	//readline.PcItem("exit"),
)