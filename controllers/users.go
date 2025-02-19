package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Add User godoc
// @Summary User
// @Description Add User
// @Tags User
// @Accept mpfd
// @Produce json
// @Security ApiKeyAuth
// @Param first_name formData string true "First Name User"
// @Param last_name formData string true "Last Name User"
// @Param phone_number formData string true "Phone Number User"
// @Param image formData file true "Image User"
// @Param email formData string true "Email User"
// @Param password formData string true "Password User"
// @Success 200 {object} Response{results=models.UserAdmin}
// @Router /users [post]
func AddUserAdmin(ctx *gin.Context) {
	var profile models.UserAdmin
	// handling body form without file
	ctx.ShouldBind(&profile)
	f, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Unauthorized file name",
		})
	}

	profile.Email = f.Value["email"][0]
	profile.Password = f.Value["password"][0]

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
		ctx.SaveUploadedFile(file, fmt.Sprintf("updload/profile/%s", storedFile))
		profile.Image = storedFile
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
	updated := models.AddUsers(profile)
	if profile.Password != "" {
		hash := lib.CreateHash(profile.Password)
		profile.Password = hash
		updated.Password = profile.Password
		updated.Email = profile.Email
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: updated,
	})
}
