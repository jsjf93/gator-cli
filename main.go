package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/jsjf93/gator-cli/internal"
	"github.com/jsjf93/gator-cli/internal/commands"
	"github.com/jsjf93/gator-cli/internal/config"
	"github.com/jsjf93/gator-cli/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	defer db.Close()

	state := internal.NewState(&cfg, dbQueries)

	cmds := commands.NewCommands()
	if err := cmds.Register("register", commands.HandlerRegister); err != nil {
		log.Fatalln(err)
	}
	if err := cmds.Register("login", commands.HandlerLogin); err != nil {
		log.Fatalln(err)
	}

	args := os.Args

	if len(args) < 3 {
		log.Fatalln("error running program. 2 args are required")
	}

	cmd := commands.Command{
		Name: args[1],
		Args: []string{args[2]},
	}

	if err := cmds.Run(&state, cmd); err != nil {
		log.Fatalln(err)
	}
}
