package user

import "errors"

var ErrEmailAlreadyUsed = errors.New("email already used")
var ErrUserNotExist = errors.New("no user with login email")
var ErrWrongPassword = errors.New("wrong password")