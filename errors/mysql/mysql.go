package mysql

import (
	"strings"

	"BookstoreUtils/errors/rest_errors"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "sql: no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return rest_errors.NewNotFoundError(err, "no record matching")
		}

		return rest_errors.NewInternalServerError(err, "error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError(err, "duplicate data")
	case 1364:
		return rest_errors.NewBadRequestError(err, "missing field")
	}

	return rest_errors.NewInternalServerError(err, "error processing request")

}
