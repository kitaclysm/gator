package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/kitaclysm/gator/internal/config"
	"github.com/kitaclysm/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbURL := cfg.DbURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s := state{db: dbQueries, cfg: cfg}

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(1)
	}
	inputCmd := os.Args[1]
	inputCrit := os.Args[2:]

	userCmd := command{
		name: inputCmd,
		args: inputCrit,
	}

	err = cmds.run(&s, userCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
