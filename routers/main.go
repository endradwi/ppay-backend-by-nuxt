package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{"GET", "POST", "PATCH"},
	}))
	ProfileRouter(router.Group("/profile"))
	AuthRouters(router.Group("/auth"))
	MovieRouter(router.Group("/movies"))
	OrdersRouter(router.Group("/orders"))
	UsersRouter(router.Group("/users"))
}
