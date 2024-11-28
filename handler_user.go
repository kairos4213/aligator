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
	printUser(user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("%v does not take any args", cmd.name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}

	for _, name := range users {
		formattedName := fmt.Sprintf("* %s", name)
		if name == s.cfg.CurrentUserName {
			formattedName = formattedName + " (current)"
		}
		fmt.Println(formattedName)
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * id:			%v\n", user.ID)
	fmt.Printf(" * name:		%v\n", user.Name)
	fmt.Printf(" * created:		%v\n", user.CreatedAt)
	fmt.Printf(" * last_updated:	%v\n", user.UpdatedAt)
}
