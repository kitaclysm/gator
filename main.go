package main

import (
	"fmt"
	"os"

	"github.com/kitaclysm/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	s := state{ cfg: cfg }

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(1)
	}
	inputCmd := os.Args[1]
	inputCrit := os.Args[2:]
	
	userCmd := command{
		name: 	inputCmd,
		args:	inputCrit,
	}

	err = cmds.run(&s, userCmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return
}