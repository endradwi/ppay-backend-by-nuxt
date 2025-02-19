package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {

	router.POST("", controllers.AddUserAdmin)
}
