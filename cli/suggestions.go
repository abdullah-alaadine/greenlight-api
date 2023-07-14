package cli

import "github.com/c-bata/go-prompt"

type suggestionType int

const (
	Commands suggestionType = iota
)

var commandSuggestions = []prompt.Suggest{
	{Text: ".quit", Description: "Quit/Exit the program"},
	{Text: ".exit", Description: "Quit/Exit the program"},
}

var suggestionsMap = map[suggestionType][]prompt.Suggest{
	Commands: commandSuggestions,
}
