package commands

import (
	"context"

	"github.com/jsjf93/gator-cli/internal"
)

func HandlerReset(state *internal.State, cmd Command) error {
	if err := state.Db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	return nil
}
