package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test/models"

	"github.com/gin-gonic/gin"
)

// Orders godoc
// @Schemes
// @Description Add Choose Cinema tikers
// @Tags Order Tiket
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Select Movie Tiket"
// @Param searchDate query string true "Search Date Cinema"
// @Param searchTime query string true "Search Time Cinema"
// @Param searchLocation query string true "Search Location Cinema"
// @Success 200 {object} Response{results=models.MoviesCinema}
// @Router /orders/cinema/{id} [get]
func OrderMovies(ctx *gin.Context) {
	var orderMovie models.OrderBody
	err := ctx.ShouldBind(&orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	order := models.OrderTicket(orderMovie)
	log.Println("data apa =", order)
	// var orderPayment models.Payment

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Order tiket sukses",
		Results: order,
	})

}

// Orders godoc
// @Schemes
// @Description Add Choose Cinema tikers
// @Tags Order Tiket
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Select Movie Tiket"
// @Param searchDate query string true "Search Date Cinema"
// @Param searchTime query string true "Search Time Cinema"
// @Param searchLocation query string true "Search Location Cinema"
// @Success 200 {object} Response{results=models.MoviesCinema}
// @Router /orders/cinema/{id} [get]
// func GetMovieCinema(ctx *gin.Context) {
// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(id)
// 	searchDate := ctx.DefaultQuery("searchDate", "")
// 	log.Println("seardate =", searchDate)
// 	searchTime := ctx.DefaultQuery("searchTime", "")
// 	log.Println("seartime =", searchTime)
// 	searchLocation := ctx.DefaultQuery("searchLocation", "")
// 	log.Println("searlocation =", searchLocation)
// 	// var cinema models.ListCinema

// 	cinema := models.BookingCinema(id, searchTime, searchDate, searchLocation)
// 	log.Print("print cinema =", cinema)
// 	// log.Println(find.Cinema_time)

// 	ctx.JSON(http.StatusOK, Response{
// 		Success: true,
// 		Message: "Cinema Selected And Choose Yout seat",
// 		Results: cinema,
// 	})
// }

func GetCinem(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var input models.MoviesCinema

	if err := ctx.ShouldBind(&input); err != nil {
		fmt.Println("Binding Error:", err)
		return
	}

	fmt.Println("input", input)
	// Query cinema
	cinema := models.BookingCinema(id, input.Location)
	// if err != nil {
	// 	return
	// }
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Cinema Selected And Choose Yout seat",
		Results: cinema,
	})
}
