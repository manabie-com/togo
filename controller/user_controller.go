package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/qgdomingo/todo-app/interfaces"
	"github.com/qgdomingo/todo-app/model"
)

type UserController struct {
	UserRepo interfaces.IUserRepository
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var user model.UserLogin

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Error on binding data from request",
			"error"   : err.Error() })
		return
	}

	if user.Username != "" && user.Password != "" {
		isUserCredentialsMatch, errMessage := uc.UserRepo.LoginUserDB(&user)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}

		if isUserCredentialsMatch {
			c.IndentedJSON(http.StatusOK, gin.H{ "message" : "Login successful" })
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{ "message" : "Entered credentials do not match" })
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Request has either empty username or password" })
	}

}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user model.NewUser

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message" : "Error on binding data from request",
			"error"   : err.Error() })
		return
	}

	if user.Username != "" && user.Name != "" && user.Email != "" && user.Password != "" && user.RepeatPassword != "" {
		if user.Password == user.RepeatPassword {
			isUserCreated, errMessage := uc.UserRepo.RegisterUserDB(&user)

			if errMessage != nil {
				c.IndentedJSON(http.StatusInternalServerError, errMessage)
				return
			}

			if isUserCreated {
				c.IndentedJSON(http.StatusOK, gin.H{ "message" : "User registration successful" })
			} else {
				c.IndentedJSON(http.StatusNotAcceptable, gin.H{ "message" : "User registration failed" })
			}

		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Passwords entered do not match" })	
		}
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Either one or more of the required data on the request is empty" })
	}

}

func (uc *UserController) FetchUserDetails(c *gin.Context) {
	userName := c.Param("username")

	if userName != "" {
		userDetailsList, errMessage := uc.UserRepo.FetchUserDetailsDB(userName)

		if errMessage != nil {
			c.IndentedJSON(http.StatusInternalServerError, errMessage)
			return
		}

		if len(userDetailsList) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "No user data was found for the specified username" })
		} else {
			c.IndentedJSON(http.StatusOK, userDetailsList)
		}

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "No username was entered on the request URL" })
	}
}

func (uc *UserController) UpdateUserDetails(c *gin.Context) {
	userName := c.Param("username")

	if userName != "" {
		var user model.UserDetails

		err := c.ShouldBindJSON(&user)
	
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message" : "Error on binding data from request",
				"error"   : err.Error() })
			return
		}

		if user.Username != "" && user.Name != "" && user.Email != "" {
			isUserUpdated, errMessage := uc.UserRepo.UpdateUserDetailsDB(&user, userName)

			if errMessage != nil {
				c.IndentedJSON(http.StatusInternalServerError, errMessage)
				return
			}
		
			if isUserUpdated {
				c.IndentedJSON(http.StatusOK, gin.H{ "message" : "User has been updated successfully" })
			} else {
				c.IndentedJSON(http.StatusNotFound, gin.H{ "message" : "User was not updated, user with the provided username is not found." })
			}

		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Either one or more of the required data on the request is empty" })
		}

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Empty username was entered on the request URL" })
	}

}

func (uc *UserController) UserPasswordChange(c *gin.Context) {
	userName := c.Param("username")

	if userName != "" {
		var user model.UserNewPassword

		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message" : "Error on binding data from request",
				"error"   : err.Error() })
			return
		}

		if user.CurrentPassword != "" && user.NewPassword != "" && user.RepeatPassword != "" {

			if user.NewPassword == user.RepeatPassword {
				isUserPasswordUpdated, errMessage := uc.UserRepo.UserPasswordChangeDB(&user, userName)

				if errMessage != nil {
					c.IndentedJSON(http.StatusInternalServerError, errMessage)
					return
				}
			
				if isUserPasswordUpdated {
					c.IndentedJSON(http.StatusOK, gin.H{ "message" : "User password has been updated successfully" })
				} else {
					c.IndentedJSON(http.StatusUnauthorized, gin.H{ "message" : "User password was not updated, current password entered does not match" })
				}

			} else {
				c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Passwords entered do not match" })	
			}

		} else {
			c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Either one or more of the required data on the request is empty" })
		}

	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message" : "Empty username was entered on the request URL" })
	}
}