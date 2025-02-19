package controllers

import (
	"log"
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Schemes
// @Description Registrasi Account
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Input Email"
// @Param password formData string true "Input Password"
// @Success 200 {object} Response
// @Router /auth/register [post]
func AuthRegister(ctx *gin.Context) {
	var form models.Users
	err := ctx.ShouldBind(&form)
	if err != nil {
		if strings.Contains(err.Error(), "Field validation for 'Email' failed on the 'email' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid Email Format",
			})
			return
		}
		if strings.Contains(err.Error(), "Field validation for 'Password' failed on the 'min' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Minimum Password 6",
			})
			return
		}
		if strings.Contains(err.Error(), "Field validation for 'Password' failed on the 'containsany' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include 1 Uppercase OR 1 Number",
			})
			return
		}

		log.Println(err)
	}
	findUser := models.FindOneUserByEmail(form.Email)
	// fmt.Println("error = ", findUser)
	if form.Email == findUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Has Registered",
		})
		return
	}
	hash := lib.CreateHash(form.Password)
	form.Password = hash
	newUser := models.InsertUser(form)

	if newUser.Id > 0 {
		profile := models.RelationProfile{
			First_Name:   "",
			Last_Name:    "",
			Image:        "",
			Phone_Number: "",
			User_Id:      newUser.Id,
		}

		models.AddProfile(profile)

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register Success"})
	} else {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "User Registration Failed",
		})
	}
}

// Auth godoc
// @Schemes
// @Description Login Account
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Input Email"
// @Param password formData string true "Input Password"
// @Success 200 {object} Response
// @Router /auth/login [post]
func AuthLogin(ctx *gin.Context) {
	var form models.Users
	ctx.ShouldBind(&form)

	foundUser := models.FindOneUserByEmail(form.Email)
	if form.Email != foundUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	match := lib.GenerateTokenArgon(form.Password, foundUser.Password)
	if !match {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	token := lib.GeneretedToken(struct {
		UserId int `json:"userId"`
	}{
		UserId: foundUser.Id,
	})
	// fmt.Println("data = ", token)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Login Success",
		Results: token,
	})

}
