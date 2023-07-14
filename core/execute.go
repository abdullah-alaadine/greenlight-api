package core

import (
	"github.com/abdullah-alaadine/sql-lite/compiler"
	"github.com/rs/zerolog/log"
)

func ExecuteStatement(statement compiler.Statement) {
	switch statement.Type {
	case compiler.Insert:
		log.Info().Msg("This is where we would like to insert")
	case compiler.Select:
		log.Info().Msg("This is where we would like to select")
	}
}
