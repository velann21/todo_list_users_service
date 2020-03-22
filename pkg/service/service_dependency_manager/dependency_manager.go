package dao_dependency_manager

import (
	"github.com/todo_list_users_service/pkg/service"
)

const(
	USERSERVICE = "UserService"
)

func NewService(objectType string)service.UserService{
	if objectType == USERSERVICE{
		return &service.UserManagementService{}
	}
	return nil
}

