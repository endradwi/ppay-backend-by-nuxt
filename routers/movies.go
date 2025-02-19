package routers

import (
	"test/controllers"
	"test/docs"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	// router.Use(middlewares.ValidationToken())

	docs.SwaggerInfo.BasePath = ""
	// router.GET("", ginSwagger.WrapHandler(swaggerfile.Handler))
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMoviesById)
	router.POST("", middlewares.ValidationToken(), controllers.SaveMovies)
	router.PATCH("/:id", controllers.EditMovie)
	router.DELETE("/:id", controllers.DeleteMovie)
}
