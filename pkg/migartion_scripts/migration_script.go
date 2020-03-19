package migartion_scripts

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
)

func MigrateDb() error {
	fmt.Println("In main")
	fsrc, err := (&file.File{}).Open("file://")
	if err != nil {
		fmt.Println(err)
		return err
	}
	db, err := sql.Open("mysql", "root:root@tcp(database)/UsersService?multiStatements=true")
	if err != nil {
		return err
	}
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithInstance("", fsrc, "usersrv", driver, )
	if err != nil{
		fmt.Println("Eror")
		return err
	}
	err = m.Steps(1)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Done")
	defer m.Close()
	return nil
}

