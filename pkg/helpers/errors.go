package helpers

import "errors"

var (
	// ErrParamMissing required field missing error
	NotValidRequestBody = errors.New("Invalid Request")
	InvalidPassword = errors.New("Invalid Password")
	ErrUserNotFound = errors.New("User Not found")
	NoresultFound = errors.New("No results found")
	InvalidRequest = errors.New("Invalid Request")
	SomethingWrong = errors.New("Something went wrong")
	UserAlreadyExist = errors.New("User with email already exists")
	Roles_notFound = errors.New("User creation failed")
)
