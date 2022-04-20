package routes

import (
	"interview/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	route := gin.Default()

	route.POST("challenge-1", controllers.TimeCalculation)
	route.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
	})

	route.GET("/movies", controllers.GetAllMovie)
	route.POST("/movies", controllers.CreateMovie)
	route.GET("/movies/:id", controllers.GetMovieById)
	route.PATCH("/movies/:id", controllers.UpdateMovie)
	route.DELETE("movies/:id", controllers.DeleteMovie)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return route
}
