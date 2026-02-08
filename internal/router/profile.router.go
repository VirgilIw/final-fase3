package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/virgilIw/final-fase3/internal/controller"
	"github.com/virgilIw/final-fase3/internal/middleware"
	"github.com/virgilIw/final-fase3/internal/repository"
	"github.com/virgilIw/final-fase3/internal/service"
)

func ProfileRouter(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	profileRouter := app.Group("/profile")

	profileRepository := repository.NewProfileRepository()
	profileService := service.NewProfileService(profileRepository, rdb, db)
	profileController := controller.NewProfileController(profileService)

	profileRouter.GET("/:id", middleware.UserOnly, profileController.GetProfile)
	profileRouter.POST("/input", middleware.UserOnly, profileController.InputProfile)
}
