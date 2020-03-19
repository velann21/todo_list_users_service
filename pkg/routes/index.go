package routes

import (
	"github.com/gorilla/mux"
	"github.com/todo_list_users_service/pkg/controller"
)

func Initialize(indexRoute *mux.Router) {
	indexRoute.HandleFunc("/signin", controller.UserSignIn).Methods("POST")
	indexRoute.HandleFunc("/signup", controller.UserSignUp).Methods("POST")
	indexRoute.HandleFunc("/getUserDetails", controller.GetUserDetails).Methods("GET")
	indexRoute.HandleFunc("/roles", controller.GetRoles).Methods("GET")
	indexRoute.HandleFunc("/createRoles", controller.CreateRoles).Methods("POST")
	indexRoute.HandleFunc("/migrateDB", controller.MigrateDB).Methods("POST")
}