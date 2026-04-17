package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jsjf93/gator-cli/internal"
)

func HandlerLogin(state *internal.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments provided")
	}

	username := cmd.Args[0]

	if len(strings.TrimSpace(username)) == 0 {
		return errors.New("invalid username provided")
	}
	if _, err := state.Db.GetUser(context.Background(), username); err != nil {
		return fmt.Errorf("unable to find user: %s", username)
	}
	if err := state.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("User: %s has been set\n", username)

	return nil
}
