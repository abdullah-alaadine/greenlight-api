package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/abdullah-alaadine/sql-lite/core"
	prompt "github.com/c-bata/go-prompt"
	"github.com/marianogappa/sqlparser"
	"github.com/rs/zerolog/log"
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
	for _, s := range suggestionsMap {
		suggestions = append(suggestions, s...)
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func getExecuter(file string) func(string) {
	// TODO: set up file for database
	return func(line string) {
		line = strings.TrimSpace(line)
		line = strings.ToLower(line)
		switch line {
		case "":
			return
		case ".exit", ".quit":
			log.Info().Msg("Bye!")
			os.Exit(0)
		default:
			// TODO: proccess commmands here
			if strings.HasPrefix(line, ".") {
				log.Error().Msg(fmt.Sprintf("unrecognized command '%s'", line))
				break
			}
			// TODO: prepare statement with SQL compiler
			query, err := sqlparser.Parse(line)
			if err == nil {
				core.ExecuteStatement(query)
			} else {
				log.Error().Msg(fmt.Sprintf("unrecognized keyword at start of '%s'", line))
			}
		}
	}
}
