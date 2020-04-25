package dao;

import (
	"context"
	"database/sql"
	"github.com/velann21/todo_list_users_service/pkg/entities"
	"github.com/velann21/todo_list_users_service/pkg/entities/data"
	req "github.com/velann21/todo_list_users_service/pkg/entities/requests"
	mysqlQuery "github.com/velann21/todo_list_users_service/pkg/helpers/mysql_query_helper"
	"golang.org/x/crypto/bcrypt"
)


type UserDaoImpl struct {
	DB *sql.DB
}

func (userDao *UserDaoImpl) CreateUserAndRoles(ctx context.Context, data entities.UserData, role []int)(int64, error){
	tx, err := userDao.DB.Begin()
	var id int64 = 0
	password, err := hashPassword(data.Password)
	if err!=nil{
		return id, err
	}

	stmt, err := mysqlQuery.UserCreateQuery(tx)
	if err!=nil{
		return id, err
	}
	res, err := stmt.Exec(data.FirstName, data.LastName, data.Email, data.PhoneNumber, data.DOB, password)
	if err != nil{
		_ = tx.Rollback()
		return id, err
	}
	id, err = res.LastInsertId()
    if err != nil{
		_ = tx.Rollback()
		return id, err
	}
	stmt, err = mysqlQuery.UserRolesUpdate(tx);
	if err != nil{
		_ = tx.Rollback()
		return id, err
	}
	for _,role := range role{
		res, err = stmt.Exec(id, role)
		if err != nil{
			_ = tx.Rollback()
			return id, err
		}

	}
	err = tx.Commit()
	if err != nil{
		return id, err
	}
	return id, nil
}


func (userDao *UserDaoImpl) GetUserByEmail(ctx context.Context, email string) ([]data.UserDataResponseWithRolePermission, error) {
	 userDataResponses := []data.UserDataResponseWithRolePermission{}
	 query := mysqlQuery.GetUserWithRole()
	 rows, err := userDao.DB.Query(query, email)
	 if err !=  nil{
		 return nil, err
	 }
	for rows.Next() {
		userDataResp := data.UserDataResponseWithRolePermission{}
		err =  rows.Scan(&userDataResp.ID, &userDataResp.FirstName,&userDataResp.LastName, &userDataResp.Email,&userDataResp.Password,&userDataResp.PhoneNumber,&userDataResp.Dob,&userDataResp.UserID,&userDataResp.RoleID,&userDataResp.RoleName)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} else {
				return nil, err
			}
		}
		userDataResponses = append(userDataResponses, userDataResp)
	}
	 return userDataResponses, nil
}

func (userDao *UserDaoImpl) GetRoles(ctx context.Context)(*sql.Rows, error){
	query := mysqlQuery.GetRoles()
	results,err := userDao.DB.Query(query)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}else{
			return nil, err
		}
	}
	return results, nil
}

func (userDao *UserDaoImpl) CreateRoles(ctx context.Context, roles req.CreateRoles) error{
	tx, err := userDao.DB.Begin()
	if err != nil{
		return err
	}
	stmt, err := mysqlQuery.CreateRoles(tx)
	if err != nil{
		_ = tx.Rollback()
		return err
	}
	_, err = stmt.Exec(roles.RoleName, roles.RoleDescription)
	if err != nil{
		_ = tx.Rollback()
		return  err
	}
	err = tx.Commit()
	if err != nil{
		_ = tx.Rollback()
		return  err
	}
	return nil
}

func (userDao *UserDaoImpl) CheckPasswordHash(ctx context.Context, password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
