package app

import (
	"github.com/superbkibbles/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)
	router.POST("/users/login", user.Login)
	router.GET("/internal/users/search", user.Search)
}
