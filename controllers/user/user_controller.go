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

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestErr("user ID should be a number")
	}
	return userID, nil
}

func Get(c *gin.Context) {
	userID, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusFound, user.Marshal(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		resErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userID, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		resErr := errors.NewBadRequestErr("invalid json body")
		c.JSON(resErr.Status, resErr)
		return
	}
	user.Id = userID

	isPartial := c.Request.Method == http.MethodPatch

	res, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, res.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userID, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if deleteErr := services.UsersService.DeleteUser(userID); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// SearchUser : get user from database
func SearchUser(c *gin.Context) {
	id := c.Param("user_id")
	c.String(http.StatusFound, fmt.Sprintf("id %s", string(id)))
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestErr("invalid Json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}
