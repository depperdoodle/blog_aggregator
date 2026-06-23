package main

import (
	"context"
	"fmt"
)

// you'll need context and fmt

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Successfully reset!")
	return nil
}
