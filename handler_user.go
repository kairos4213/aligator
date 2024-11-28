package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kairos4213/aligator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes <name> arg", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Printf("User %v logged in\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes <name> arg", cmd.name)
	}

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set user: %w", err)
	}

	fmt.Println("User has been created!")
	fmt.Printf("id: %v, time_created: %v, last_updated: %v, name: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
