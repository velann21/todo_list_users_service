package service

import (
	"context"
	"github.com/todo_list_users_service/pkg/entities/data"
	dataResponse "github.com/todo_list_users_service/pkg/entities/data"
	requestEntites "github.com/todo_list_users_service/pkg/entities/requests"
)
type UserService interface {
	UserSignIn(ctx context.Context, request requestEntites.UserSigninRequest) (*string, error);
	UserSignUp(ctx context.Context, request requestEntites.UserSignupRequest) (*string, error);
	GetUser(ctx context.Context, request requestEntites.GetUserDetails) ([]data.UserDataResponseWithRolePermission,error);
	GetRoles(ctx context.Context) ([]dataResponse.UserRolesResponse, error);
	CreateRoles(ctx context.Context, roles requestEntites.CreateRoles) error
	MigrateDBService(ctx context.Context) error
}
