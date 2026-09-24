package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/warmdev17/hapi/config"
	"github.com/warmdev17/hapi/internal/auth/application"
	authtransport "github.com/warmdev17/hapi/internal/auth/transport"
	db "github.com/warmdev17/hapi/internal/db/sqlc"
	"github.com/warmdev17/hapi/internal/infrastructure/jwt"
	"github.com/warmdev17/hapi/internal/infrastructure/password"
	userinfra "github.com/warmdev17/hapi/internal/user/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

type Module struct {
	Handler *authtransport.Handler
}

func NewModule(
	queries *db.Queries,
	cfg *config.Config,
	api *gin.RouterGroup,
) *Module {
	userRepo := userinfra.NewPostgresUserRepository(queries)
	passwordHasher := password.NewBcryptHasher(bcrypt.DefaultCost)
	tokenProvider := jwt.NewJWTProvider(cfg.JWT.Secret)

	registerService := application.NewRegisterService(userRepo, passwordHasher, tokenProvider, cfg.JWT.AccessTTL)
	handler := authtransport.NewHandler(registerService)

	auth := api.Group("/auth")
	authtransport.RegisterRoutes(auth, handler)

	return &Module{
		Handler: handler,
	}
}
