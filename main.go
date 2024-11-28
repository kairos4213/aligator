package main

import (
	"fmt"

	"github.com/kairos4213/aligator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("james")
	cfg = config.Read()
	fmt.Println(cfg)
}
