package routes

import (
	"github.com/gorilla/mux"
	"github.com/velann21/todo_list_users_service/pkg/controller"
	"github.com/velann21/todo_list_users_service/pkg/dao"
	"github.com/velann21/todo_list_users_service/pkg/databases"
	"github.com/velann21/todo_list_users_service/pkg/service"
)

func Initialize(indexRoute *mux.Router) {
	indexRoute.HandleFunc("/signin", newController().UserSignIn).Methods("POST")
	indexRoute.HandleFunc("/signup", newController().UserSignUp).Methods("POST")
	indexRoute.HandleFunc("/getUserDetails", newController().GetUserDetails).Methods("GET")
	indexRoute.HandleFunc("/roles", newController().GetRoles).Methods("GET")
	indexRoute.HandleFunc("/createRoles", newController().CreateRoles).Methods("POST")
	indexRoute.HandleFunc("/migrateDB", newController().MigrateDB).Methods("POST")
}

func newController() controller.Controller{
	controllerObj := controller.Controller{Service: &service.UserManagementService{Dao: &dao.UserDaoImpl{databases.GetSqlConnection()}}}
	return controllerObj
}