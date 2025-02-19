package middlewares

import (
	"net/http"
	"os"
	"strings"
	"test/controllers"
	"test/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/joho/godotenv"
)

func ValidationToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.Header("Content-Type")
		head := ctx.GetHeader("Authorization")

		if head == "" {
			ctx.JSON(http.StatusNotFound, controllers.Response{
				Success: false,
				Message: "Token not found",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]

		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
		// log.Println("data =", tok)
		// out := jwt.Claims{}
		out := make(map[string]interface{})

		godotenv.Load()
		err := tok.Claims([]byte(lib.GetMD5hash(os.Getenv("JWT_SECRET"))), &out)

		ctx.Set("userId", int(out["userId"].(float64)))

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.Response{
				Success: false,
				Message: "Unauthorized",
			})

			ctx.Abort()
		}
		// if head != "true" {
		// 	ctx.JSON(http.StatusUnauthorized, controllers.Response{
		// 		Success: false,
		// 		Message: "Unauthorized",
		// 	})
		// 	ctx.Abort()
		// }
		ctx.Next()
	}
}
