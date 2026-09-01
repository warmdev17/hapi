package main

import (
	"context"
	"fmt"

	"github.com/warmdev17/hapi/config"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()
	fmt.Println(ctx, cfg)
}
