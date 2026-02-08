package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/virgilIw/final-fase3/internal/middleware"
)

func Init(app *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	app.Use(middleware.CORSMiddleware())
	AuthRouter(app, db, rdb)
	ProfileRouter(app, db, rdb)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
