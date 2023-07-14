package cli

import (
	prompt "github.com/c-bata/go-prompt"
)

func StartPrompt(file string) {
	p := prompt.New(
		getExecuter(file),
		completer,
		prompt.OptionTitle("Welcome to SQL Lite"),
		prompt.OptionPrefix("sql > "),
	)
	p.Run()
}

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
