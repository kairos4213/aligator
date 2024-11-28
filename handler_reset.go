package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("%v does not take any args", cmd.name)
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users table data: %w", err)
	}

	fmt.Println("Database reset successful")
	return nil
}
