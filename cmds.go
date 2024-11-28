package main

import "errors"

type commands struct {
	registeredCmds map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	callbackf, ok := c.registeredCmds[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}

	return callbackf(s, cmd)
}
