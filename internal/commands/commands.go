package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jsjf93/gator-cli/internal"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	registry map[string]func(*internal.State, Command) error
}

func NewCommands() Commands {
	return Commands{
		registry: make(map[string]func(*internal.State, Command) error),
	}
}

func (c *Commands) Run(state *internal.State, cmd Command) error {
	fn, found := c.registry[cmd.Name]
	if !found {
		return fmt.Errorf("no matching commands with name: %s", cmd.Name)
	}

	err := fn(state, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *Commands) Register(name string, f func(*internal.State, Command) error) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("a command requires a name")
	}

	if _, found := c.registry[name]; found {
		return fmt.Errorf("command: %s is already registered", name)
	}

	c.registry[name] = f

	return nil
}
