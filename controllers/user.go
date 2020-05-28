package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tesh254/golang_todo_api/forms"
	"github.com/tesh254/golang_todo_api/helpers"
	"github.com/tesh254/golang_todo_api/models"
	"github.com/tesh254/golang_todo_api/services"
)

// Import the userModel from the models
var userModel = new(models.UserModel)

// UserController defines the user controller methods
type UserController struct{}

// Signup controller handles registering a user
func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand

	// Bind the data from the request body to the SignupUserCommand Struct
	// Also check if all fields are provided
	if c.BindJSON(&data) != nil {
		// specified response
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		c.Abort()
		// return nothing
		return
	}

	/*
		You can add your validation logic
		here such as email

		if regexMethodChecker(data.Email) {
			c.JSON(400, gin.H{"message": "Email is invalid"})
			c.Abort()
			return
		}
	*/
	result, _ := userModel.GetUserByEmail(data.Email)

	// If there happens to be a result respond with a
	// descriptive mesage
	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err := userModel.Signup(data)

	// Check if there was an error when saving user
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}

// Login allows a user to login a user and get
// access token
func (u *UserController) Login(c *gin.Context) {
	var data forms.LoginUserCommand

	// Bind the request body data to var data and check if all details are provided
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	// Get the hashed password from the saved document
	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		c.Abort()
		return
	}

	jwtToken, err2 := services.GenerateToken(data.Email)

	// If we fail to generate token for access
	if err2 != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})
}

// PasswordReset handles user password request
func (u *UserController) PasswordReset(c *gin.Context) {
	var data forms.PasswordResetCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	if (data.Password != data.Confirm) {
		c.JSON(400, gin.H{"message": "Passwords do not match"})
		c.Abort()
		return
	}

	resetToken, _ := c.GetQuery("reset_token")

	userID, _ := services.DecodeInfiniteToken(resetToken)

	result, err := userModel.GetUserByID(userID)

	if err != nil {
		// Return response when we get an error while fetching user
		c.JSON(500, gin.H{"message": "Something wrong happened, try again later"})
		c.Abort()
		return
	}
	
	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User accoun was not found"})
		c.Abort()
		return
	}
	// Hash the new password
	newHashedPassword := helpers.GeneratePasswordHash([]byte(data.Password))

	// Update user account
	_err := userModel.UpdateUserPass(userID, newHashedPassword)

	if _err != nil {
		// Return response if we are not able to update user password
		c.JSON(500, gin.H{"message": "Somehting happened while updating your password try again"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "Password has been updated, log in"})
	c.Abort()
	return
}