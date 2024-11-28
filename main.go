package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kairos4213/aligator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	ste := state{cfg: &cfg}
	cmds := commands{
		registeredCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Printf("must provide at least one argument\n")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	if err = cmds.run(&ste, cmd); err != nil {
		log.Fatal(err)
	}
}
