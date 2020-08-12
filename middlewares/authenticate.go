package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tesh254/golang_todo_api/models"
	"github.com/tesh254/golang_todo_api/services"
)

func responseWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

// Authenticate fetches user details from token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		requiredToken := c.Request.Header["Authorization"]

		if len(requiredToken) == 0 {
			responseWithError(c, 403, "Please login to your account")
		}

		userID, _ := services.DecodeToken(requiredToken[0])

		result, err := new(models.UserModel).GetUserByEmail(userID)

		if result.Email == "" {
			responseWithError(c, 404, "User account not found")
			return
		}

		if err != nil {
			responseWithError(c, 500, "Something went wrong giving you access")
			return
		}

		c.Set("User", result)

		c.Next()
	}
}
