package users

import (
	"fmt"
	"strings"

	"github.com/superbkibbles/bookstore_users-api/datasourses/mysql/users_db"
	"github.com/superbkibbles/bookstore_users-api/logger"
	"github.com/superbkibbles/bookstore_users-api/utils/errors"
	"github.com/superbkibbles/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUsers          = "INSERT INTO users(first_name, last_name, email, date_created, status, password) values(?,?,?,?,?,?);"
	queryGetUser              = "SELECT id, first_name, last_name, email, date_created, status FROM users where id=?;"
	queryUpdateUser           = "UPDATE users SET first_name=?, last_name=?, email=? where id=?;"
	queryDeleteUser           = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus     = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status from users WHERE email=? AND password=? AND status=?"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying to prepare get user statment", err)
		return errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error while trying to get user by ID", getErr)
		return errors.NewInternalServerErr("Database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUsers)
	if err != nil {
		logger.Error("error while trying to prepare save user statment", err)
		return errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()

	insertRes, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error while trying to save user", err)
		return errors.NewInternalServerErr("Database error")
	}

	userID, err := insertRes.LastInsertId()
	if err != nil {
		logger.Error("error while trying to get created user id", err)
		return errors.NewInternalServerErr("Database error")
	}
	user.Id = userID

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while trying to prepare update user statment", err)
		return errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.LastName, user.Id)
	if err != nil {
		logger.Error("error while trying to update user", err)
		return errors.NewInternalServerErr("Database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error while trying to prepare Delete user statment", err)
		return errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error while trying to delere user", err)
		return errors.NewInternalServerErr("Database error")
	}

	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error while trying to prepare find user by status statment", err)
		return nil, errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error while trying to find user by status", err)
		return nil, errors.NewInternalServerErr("Database error")
	}
	defer rows.Close()
	resutls := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error while trying to scan user into user struct", err)
			return nil, errors.NewInternalServerErr("Database error")
		}
		resutls = append(resutls, user)
	}

	if len(resutls) == 0 {
		return nil, errors.NewNotFoundErr(fmt.Sprintf("No user matching status %s", status))
	}
	return resutls, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindEmailAndPassword)
	if err != nil {
		logger.Error("error while trying to prepare get user by email and password statment", err)
		return errors.NewInternalServerErr("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRow) {
			return errors.NewNotFoundErr("Invalid user credentials")
		}
		logger.Error("error while trying to get user by email and password", getErr)
		return errors.NewInternalServerErr("Database error")
	}

	return nil
}
