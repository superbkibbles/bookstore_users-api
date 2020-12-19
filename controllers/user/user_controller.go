package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superbkibbles/bookstore_users-api/domain/users"
	"github.com/superbkibbles/bookstore_users-api/services"
	"github.com/superbkibbles/bookstore_users-api/utils/errors"
)

//  Entery point of application

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestErr("user ID should be a number")
		c.JSON(err.Status, err)
		return
	}
	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusFound, result)
}

func CreateUser(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		resErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// SearchUser : get user from database
func SearchUser(c *gin.Context) {
	id := c.Param("user_id")
	c.String(http.StatusFound, fmt.Sprintf("id %s", string(id)))
}
