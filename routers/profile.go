package routers

import (
	"test/controllers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidationToken())
	router.PATCH("", controllers.EditProfile)
	router.DELETE("", controllers.DeletedProfile)
	// router.PATCH("/:id", controllers.EditUser)
	router.GET("", controllers.GetProfile)

}
