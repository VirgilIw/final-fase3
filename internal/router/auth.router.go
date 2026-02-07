package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/final-fase3/internal/controller"
	"github.com/virgilIw/final-fase3/internal/repository"
	"github.com/virgilIw/final-fase3/internal/service"
)

func AuthRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	authRouter := app.Group("/auth")

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository, rdb, db)
	authController := controller.NewAuthController(authService)
	authRouter.POST("/", authController.Login)
	authRouter.POST("/new", authController.Register)
}
