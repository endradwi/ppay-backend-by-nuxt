package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRouter(router *gin.RouterGroup) {
	// router.GET("", controllers.GetAllMovies)
	router.GET("/cinema", controllers.GetCinem)
	router.POST("/payment", controllers.ChoosePayment)
	router.POST("/payment/paid", controllers.PaidPayment)
	router.POST("", controllers.OrderMovies)
}
