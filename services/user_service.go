package services

import (
	"github.com/superbkibbles/bookstore_users-api/domain/users"
	"github.com/superbkibbles/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userID}

	if err := user.Get(); err != nil {
		return nil, err
	}
	// if userID <= 0 {
	// 	return nil, errors.NewBadRequestErr("Invalid user id")
	// }

	return user, nil
}
