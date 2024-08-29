package domain

import "errors"

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrInvalidPassword          = errors.New("invalid password")
	ErrDatabaseConnectionFailed = errors.New("database connection failed")
	ErrIDNotFound               = errors.New("id not found")
	ErrGetUserByData            = errors.New("error getting user by data")
	ErrFindUser                 = errors.New("error user not found")
	ErrToCreateUser             = errors.New("error creating user")
)

type Errors struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
