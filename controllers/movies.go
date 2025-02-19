package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// All Movies godoc
// @Summary Movies
// @Description  Get All Movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param search query string false "Search Movie"
// @Param page query int false "Page Movie"
// @Param limit query int false "Limit Movie"
// @Param sort query string false "Sort Movie"
// @Success 200 {object} Response{results=models.ListAllMovie}
// @Router /movies [get]
func GetAllMovies(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	sortmovie := ctx.DefaultQuery("sort", "ASC")
	if sortmovie != "ASC" {
		sortmovie = "DESC"
	}

	var movies models.ListAllMovie
	var count int
	get := lib.Redis().Get(context.Background(), ctx.Request.RequestURI)
	getCount := lib.Redis().Get(context.Background(),
		fmt.Sprintf("count+%s", ctx.Request.RequestURI))
	if get.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &movies)
	} else {
		movies = models.FindAllMovie(page, limit, search, sortmovie)
		change, _ := json.Marshal(movies)
		lib.Redis().Set(
			context.Background(),
			ctx.Request.RequestURI,
			string(change),
			0,
		)
	}
	if getCount.Val() != "" {
		byt := []byte(get.Val())
		json.Unmarshal(byt, &count)
	} else {
		count = models.CountData(search)
		change, _ := json.Marshal(count)
		lib.Redis().Set(context.Background(),
			fmt.Sprintf("count+%s", ctx.Request.RequestURI),
			string(change),
			0,
		)
	}
	// count = models.CountData(search)
	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	// log.Println("errorrrr", totalPage)
	// log.Println("errorrrr", count)
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	prevPage := page - 1
	if prevPage < 2 {
		prevPage = 0
	}
	ctx.JSON(200, Response{
		Success: true,
		Message: "View All Movie",
		PageInfo: PageInfo{
			CurentPage: page,
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPage:  totalPage,
			TotalData:  count,
		},
		Results: movies,
	})
}

// Detail Movie godoc
// @Schemes
// @Description Detail Movies
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Detail Movie"
// @Success 200 {object} Response{results=models.MoviesNoTag}
// @Router /movies/{id} [get]
func GetMoviesById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	find := models.FindOneMovie(id)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Detail Movie",
		Results: find,
	})

}

// Edit Movie godoc
// @Schemes
// @Description Edit Movies
// @Tags Movies
// @Accept mpfd
// @Produce json
// @Param id path int true "Id Movie"
// @Param tittle formData string true "Edit tittle"
// @Param genre formData string true "Update genre"
// @Param image formData file true "Update Image"
// @Param synopsis formData string true "Update Synopsis"
// @Param author formData string true "Update Author"
// @Param actors formData string true "Update Actors"
// @Param release_date formData string true "Update Realease Date"
// @Param duration formData string true "Update Duration"
// @Param tag formData string true "Update Tag"
// @Success 200 {object} Response{results=models.Movie_body}
// @Router /movies/{id} [patch]
func EditMovie(ctx *gin.Context) {
	paramId, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.FindOneMovie(paramId)
	// if paramId != profile.Id {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// 	})
	// 	return
	// }
	// handling body form without file

	var updatedMovie models.Movie_body
	if err := ctx.ShouldBind(&movie); err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	// f, _ := ctx.MultipartForm()
	file, _ := ctx.FormFile("image")

	if file.Filename != "" {
		filename := uuid.New().String()

		// handling extentioin .jpg dll
		splitedFilename := strings.Split(file.Filename, ".")
		ext := splitedFilename[len(splitedFilename)-1]
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Must Fill .jpg, .jpeg, .png",
			})
			return
		}

		// handling name file
		ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", storedFile))
		movie.Image = storedFile
	}

	// Validation Size File
	maxfile := 1 * 1024 * 1024
	if file.Size > int64(maxfile) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "File to Large",
		})
		return
	}

	updated := models.UpdateMovie(updatedMovie)

	fmt.Println("data upload =", updated)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update Movie Success",
		Results: updated,
	})
}

// Add Movie godoc
// @Schemes
// @Description Add Movies
// @Tags Movies
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param tittle formData string true "Edit tittle"
// @Param genre formData string true "Update genre"
// @Param image formData file true "Update Image"
// @Param synopsis formData string true "Update Synopsis"
// @Param author formData string true "Update Author"
// @Param actors formData string true "Update Actors"
// @Param release_date formData string true "Update Realease Date"
// @Param duration formData string true "Update Duration"
// @Param tag formData string true "Update Tag"
// @Success 200 {object} Response{results=models.MoviesbyTag}
// @Router /movies [post]
func SaveMovies(ctx *gin.Context) {
	var formData models.Movie_body
	text := ctx.ShouldBind(&formData)
	log.Println("data=", text)
	f, _ := ctx.MultipartForm()
	file, _ := ctx.FormFile("image")

	formData.Release_date = f.Value["release_date"][0]
	formData.Duration = f.Value["duration"][0]

	if file.Filename != "" {
		filename := uuid.New().String()

		// handling extentioin .jpg dll
		splitedFilename := strings.Split(file.Filename, ".")
		ext := splitedFilename[len(splitedFilename)-1]
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Must Fill .jpg, .jpeg, .png",
			})
			return
		}

		// handling name file
		ctx.SaveUploadedFile(file, fmt.Sprintf("upload/movies/%s", storedFile))
		formData.Image = storedFile
	}

	// Validation Size File
	maxfile := 1 * 1024 * 1024
	if file.Size > int64(maxfile) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "File to Large",
		})
		return
	}

	temp, _ := models.InsertMovie(formData)

	ctx.JSON(200, Response{
		Success: true,
		Message: "Your Movie Saved",
		Results: temp,
	})
}

// Delete Movie godoc
// @Schemes
// @Description Delete Movies
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Delete Movie"
// @Success 200 {object} Response{results=models.MoviesbyTag}
// @Router /movies/{id} [delete]
func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	deleted := models.DeleteMovie(id)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Deleted Success",
		Results: deleted,
	})

}
