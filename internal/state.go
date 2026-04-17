package internal

import (
	"github.com/jsjf93/gator-cli/internal/config"
	"github.com/jsjf93/gator-cli/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}

func NewState(config *config.Config, db *database.Queries) State {
	return State{
		Config: config,
		Db:     db,
	}
}
