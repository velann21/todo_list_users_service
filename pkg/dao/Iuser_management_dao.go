package dao

import (
	"context"
	"database/sql"
	"github.com/todo_list_users_service/pkg/entities"
	"github.com/todo_list_users_service/pkg/entities/data"
	req "github.com/todo_list_users_service/pkg/entities/requests"
)

type UserDao interface {
	CreateUserAndRoles(ctx context.Context, data entities.UserData, role []int)(int64, error);
	GetUserByEmail(ctx context.Context, email string) ([]data.UserDataResponseWithRolePermission, error)
	GetRoles(ctx context.Context)(*sql.Rows, error)
	CreateRoles(ctx context.Context, roles req.CreateRoles) error
	CheckPasswordHash(ctx context.Context, password, hash string) bool;
}
