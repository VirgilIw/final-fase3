package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/final-fase3/internal/controller"
	"github.com/virgilIw/final-fase3/internal/repository"
	"github.com/virgilIw/final-fase3/internal/service"
	"github.com/virgilIw/final-fase3/pkg/hash"
)

func AuthRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRouter := app.Group("/auth")

	hashConfig := hash.NewHashConfig(
		64*1024, // memory (64 MB)
		3,       // iterations
		2,       // parallelism
		16,      // salt length
		32,      // key length
	)

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, rdb, db, hashConfig)
	authController := controller.NewAuthController(authService)
	authRouter.POST("login", authController.Login)
	authRouter.POST("register", authController.Register)
}
