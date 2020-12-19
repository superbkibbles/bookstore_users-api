package app

import (
	"github.com/superbkibbles/bookstore_users-api/controllers/ping"
	"github.com/superbkibbles/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
	// router.POST("/user/search", user.SearchUser)
}
