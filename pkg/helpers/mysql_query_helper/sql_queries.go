package mysql_query_helper

import (
	"database/sql"
)

func UserCreateQuery(db *sql.Tx) (*sql.Stmt, error){
	res, err := db.Prepare("INSERT INTO users(first_name , last_name, email, phone_number, dob, password) VALUES( ?, ?, ?, ?,?,? )")
	return res, err
}

func GetUserByEmail()(query string){
	sqlStatement := `Select first_name,last_name,dob,phone_number,email,password from Users where email=?`
	return sqlStatement
}

func GetUserWithRole()(query string){
	sqlStatement := `Select a.id, a.first_name, a.last_name, a.email, a.password, a.phone_number, a.dob, b.user_id, b.role_id, roles.role_name from users a  Inner Join users_roles b on a.id = b.user_id Inner Join (Select id,role_name from roles) roles on b.role_id = roles.id where a.email=?;`
	return sqlStatement
}

func UserRolesUpdate(db *sql.Tx)(*sql.Stmt, error){
	stmt, err := db.Prepare("INSERT INTO users_roles(user_id , role_id) VALUES( ?, ?)")
	return stmt, err
}

func UserRoleUpdate()(string){
	sqlStatement := `INSERT INTO users_roles(user_id , role_id) VALUES($1, $2)`
	return sqlStatement
}

func GetRoles()string{
	sqlStatement := `Select id, role_name, role_description from roles;`
	return sqlStatement
}

func CreateRoles(db *sql.Tx) (*sql.Stmt, error){
	res, err := db.Prepare("INSERT INTO roles(role_name,role_description) VALUES(?, ?)")
	return res, err
}