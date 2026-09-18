package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/warmdev17/hapi/config"
	authapplication "github.com/warmdev17/hapi/internal/auth/application"
	authtransport "github.com/warmdev17/hapi/internal/auth/transport"
	db "github.com/warmdev17/hapi/internal/db/sqlc"
	"github.com/warmdev17/hapi/internal/infrastructure/jwt"
	"github.com/warmdev17/hapi/internal/infrastructure/password"
	"github.com/warmdev17/hapi/internal/infrastructure/postgres"
	userinfra "github.com/warmdev17/hapi/internal/user/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%v:%v@localhost:%v/%v?sslmode=disable", dbCfg.User, dbCfg.Password, dbCfg.Port, dbCfg.Name)
	pool, err := postgres.NewPool(ctx, dsn)
	if err != nil {
		log.Fatalf("server failed")
	}
	defer pool.Close()

	queries := db.New(pool)

	userRepo := userinfra.NewPostgresUserRepository(queries)

	passwordHasher := password.NewBcryptHasher(bcrypt.DefaultCost)
	tokenProvider := jwt.NewJWTProvider(cfg.JWT.Secret)

	registerService := authapplication.NewRegisterService(userRepo, passwordHasher, tokenProvider, cfg.JWT.AccessTTL)

	authHandler := authtransport.NewHandler(registerService)

	router := gin.Default()

	api := router.Group("/api/v1")
	auth := api.Group("/auth")

	authtransport.RegisterRoutes(auth, authHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
