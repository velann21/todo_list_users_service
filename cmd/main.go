package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_users_service/pkg/databases"
	"github.com/todo_list_users_service/pkg/helpers"
	"github.com/todo_list_users_service/pkg/routes"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:time.RFC3339,
	})
	helpers.SetEnv()

	r := mux.NewRouter().StrictSlash(false)
	mainRoutes := r.PathPrefix("/api/v1/users").Subrouter()
	routes.Initialize(mainRoutes)
	sqlconn := databases.SQLConnection{}
	err := sqlconn.OpenSqlConnection()
	if err != nil{
		logrus.WithField("EventType", "DbConnection").WithError(err).Error("Db Connection Error")
		os.Exit(100)
	}
	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : "+"8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
		return
	}
}
