package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/superbkibbles/bookstore_users-api/utils/errors"
)

const (
	ErrorNoRow = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRow) {
			return errors.NewNotFoundErr("No record matching the given id")
		}
		return errors.NewInternalServerErr("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestErr("invalid data")
	}
	return errors.NewInternalServerErr("error processing request")
}
