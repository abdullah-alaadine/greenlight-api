package main

import (
	"os"
	"time"

	"github.com/abdullah-alaadine/sql-lite/cli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	log.Info().Msg("Hello, world!")
	cli.StartPrompt("randomName")
}
