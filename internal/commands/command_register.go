package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jsjf93/gator-cli/internal"
	"github.com/jsjf93/gator-cli/internal/database"
)

func HandlerRegister(state *internal.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("no arguments provided")
	}

	username := cmd.Args[0]

	if len(strings.TrimSpace(username)) == 0 {
		return errors.New("invalid username provided")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := state.Db.CreateUser(context.Background(), params)
	if err != nil {
		return errors.New("user already exists")
	}

	state.Config.SetUser(user.Name)

	fmt.Printf("User: %s was created", user.Name)

	return nil
}
