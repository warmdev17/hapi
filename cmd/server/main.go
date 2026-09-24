package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/warmdev17/hapi/config"
	"github.com/warmdev17/hapi/internal/auth"
	db "github.com/warmdev17/hapi/internal/db/sqlc"
	"github.com/warmdev17/hapi/internal/infrastructure/postgres"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%v:%v@localhost:%v/%v?sslmode=disable", dbCfg.User, dbCfg.Password, dbCfg.Port, dbCfg.Name)
	pool, err := postgres.NewPool(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer pool.Close()

	queries := db.New(pool)

	router := gin.Default()
	api := router.Group("/api/v1")

	auth.NewModule(queries, cfg, api)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
