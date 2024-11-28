package main

import (
	"github.com/kairos4213/aligator/internal/config"
	"github.com/kairos4213/aligator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
