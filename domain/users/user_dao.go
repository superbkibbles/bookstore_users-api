package users

import (
	"fmt"

	"github.com/superbkibbles/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundErr(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d already exist", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
