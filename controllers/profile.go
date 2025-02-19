package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Profile godoc
// @Schemes
// @Description Update Movies
// @Tags Profile
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param first_name formData string true "Update First Name"
// @Param last_name formData string true "Update Last Name"
// @Param phone_number formData string true "Update Phone_Number"
// @Param image formData file false "Update Image"
// @Param email formData string false "Update Email"
// @Param password formData string false "Update Password"
// @Success 200 {object} Response{results=models.Profile}
// @Router /profile [patch]
func EditProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "Unauthorized id",
		})
		return
	}

	var profile models.Profile
	// handling body form without file
	err := ctx.ShouldBind(&profile)
	if err != nil {
		fmt.Println("Error From Body", err)
	}
	fmt.Println("data profile=", profile)

	// f, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("image")
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "No file uploaded",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Internal Server Error",
			})
		}
		return
	}

	// profile.Email = f.Value["email"][0]
	// profile.Password = f.Value["password"][0]

	if file != nil && file.Filename != "" {
		filename := uuid.New().String()

		// handling extension .jpg dll
		splittedFilename := strings.Split(file.Filename, ".")
		ext := splittedFilename[len(splittedFilename)-1]
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Must Fill .jpg, .jpeg, .png",
			})
			return
		}

		// handling nama file
		ctx.SaveUploadedFile(file, fmt.Sprintf("upload/profile/%s", storedFile))
		profile.Image = storedFile
	}

	// Validation Size File
	maxfile := 1 * 1024 * 1024
	if file != nil && file.Size > int64(maxfile) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "File is too Large",
		})
		return
	}
	if profile.Password != "" {
		hash := lib.CreateHash(profile.Password)
		profile.Password = hash
	}
	updated := models.UpdatedProfile(profile, val.(int))

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: updated,
	})
}

// Profile godoc
// @Schemes
// @Description  Get Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} Response{results=models.Profile}
// @Security ApiKeyAuth
// @Router /profile [get]
func GetProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")

	profile := models.FindOneProfile(val.(int))

	if isAvail {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Detail Profile",
			Results: profile,
		})
	}
}

// Delete profile godoc
// @Summary Profile
// @Description Delete profiles
// @Tags Profile
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Response{results=models.Profile}
// @Router /profile [delete]
func DeletedProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	deleted := models.DeleteProfile(val.(int))

	if isAvail {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Deleted Profile Success",
			Results: deleted,
		})
	}
}
