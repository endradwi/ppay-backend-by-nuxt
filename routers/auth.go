package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(router *gin.RouterGroup) {
	router.POST("/register", controllers.AuthRegister)
	router.POST("/login", controllers.AuthLogin)
}
