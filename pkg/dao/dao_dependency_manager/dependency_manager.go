package dao_dependency_manager


import "github.com/velann21/todo_list_users_service/pkg/dao"

const(
	USERDAO = "UserDao"
)

func NewDao(objectType string)dao.UserDao{
	if objectType == USERDAO{
		return &dao.UserDaoImpl{}
	}
	return nil
}

