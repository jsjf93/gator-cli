package commands

import (
	"context"
	"fmt"

	"github.com/jsjf93/gator-cli/internal"
)

func HandlerUsers(state *internal.State, cmd Command) error {
	currentUser := state.Config.CurrentUserName
	users, err := state.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		marker := ""
		if user.Name == currentUser {
			marker = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, marker)
	}

	return nil
}
