package services

import (
	"github.com/superbkibbles/bookstore_users-api/domain/users"
	"github.com/superbkibbles/bookstore_users-api/utils/crypto_utils"
	"github.com/superbkibbles/bookstore_users-api/utils/date_utils"
	"github.com/superbkibbles/bookstore_utils-go/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	GetUser(int64) (*users.User, rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(int64) rest_errors.RestErr
	SearchUser(string) (users.Users, rest_errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.RestErr)
}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowDBFromat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) GetUser(userID int64) (*users.User, rest_errors.RestErr) {
	user := &users.User{Id: userID}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) UpdateUser(partial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if partial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *usersService) DeleteUser(userID int64) rest_errors.RestErr {
	user := &users.User{Id: userID}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
