package controllers

import (
	"fmt"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

func ChoosePayment(ctx *gin.Context) {
	var paymentorder models.PaymentData
	err := ctx.ShouldBind(&paymentorder)

	if err != nil {
		fmt.Println(err)
		return
	}

	order := models.PaymentMethod(paymentorder)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Pay your order",
		Results: order,
	})
}

func PaidPayment(ctx *gin.Context) {
	var paymentorder models.PaymentPaidOrders
	err := ctx.ShouldBind(&paymentorder)

	if err != nil {
		fmt.Println(err)
		return
	}

	// order := models.PaymentPaid(paymentorder)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Payment Success",
		// Results: order,
	})
}
