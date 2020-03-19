package service;

import (
	"context"
	dm "github.com/todo_list_users_service/pkg/dao/dao_dependency_manager"
	internalRequest "github.com/todo_list_users_service/pkg/entities"
	"github.com/todo_list_users_service/pkg/entities/data"
	dataResponse "github.com/todo_list_users_service/pkg/entities/data"
	ipcRequests "github.com/todo_list_users_service/pkg/entities/ipc_requests"
	ipcResponse "github.com/todo_list_users_service/pkg/entities/ipc_responses"
	requestEntites "github.com/todo_list_users_service/pkg/entities/requests"
	"github.com/todo_list_users_service/pkg/helpers"
	migrateDb "github.com/todo_list_users_service/pkg/migartion_scripts"
	"log"
)
type UserManagementService struct{
     HttpHandler func()
}

func (um *UserManagementService) UserSignIn(ctx context.Context, request requestEntites.UserSigninRequest) (*string, error) {
	userDao := dm.NewDao(dm.USERDAO)
	userModel := internalRequest.UserData{}
	userData, err := userDao.GetUserByEmail(ctx, request.Email)
	if err != nil{
		return nil, err
	}
	if len(userData) > 0{
		return nil, helpers.UserAlreadyExist
	}
	userModel.PopulateUserData(request.FirstName, request.LastName, request.Email, request.DOB, request.PhoneNumber, request.Password);
    _ ,err = userDao.CreateUserAndRoles(ctx, userModel, request.Role)
    if err != nil{
		return nil, helpers.Roles_notFound
	}
    return &request.Email, nil
}


func (um *UserManagementService) UserSignUp(ctx context.Context, request requestEntites.UserSignupRequest) (*string, error) {
	userDao := dm.NewDao(dm.USERDAO)
	userDatas, err := userDao.GetUserByEmail(ctx, request.Email)
	if err != nil{
		return nil, err
	}
	if userDatas == nil{
		return nil, helpers.ErrUserNotFound
	}
	isPasswordMatch := userDao.CheckPasswordHash(ctx, request.Password, userDatas[0].Password)
	if isPasswordMatch != true{
		return nil, helpers.InvalidPassword
	}

	authRequestData := ipcRequests.AuthRequestBody{}
	roles := make([]int, 0)
    for _,value := range userDatas{
		roles = append(roles, value.RoleID)
	}
	authRequestData.RoleID = roles
	authRequestData.FirstName = userDatas[0].FirstName
	authRequestData.Email = userDatas[0].Email
	jsonData, err:=authRequestData.MarshalAuthRequestBody()
	log.Print(string(jsonData))
	resp, err := helpers.HttpRequest("POST",helpers.ReadEnv(helpers.AUTHSERVICECONNECTION),jsonData)
	if err != nil{
		return nil, helpers.SomethingWrong
	}
	body := resp.Body
	newTokenResp := ipcResponse.NewTokenResponse{}
	response, err := newTokenResp.PopulateNewTokenResp(body)
	if err != nil{
		return nil, err
	}

	return &response.Data[0].Token , nil
}

func (um *UserManagementService) GetUser(ctx context.Context, request requestEntites.GetUserDetails) ([]data.UserDataResponseWithRolePermission,error){
	userDao := dm.NewDao(dm.USERDAO)
	userDatas, err := userDao.GetUserByEmail(ctx, request.EmailID)
	if err != nil{
		return nil, err
	}
	if userDatas == nil{
		return nil, helpers.InvalidRequest
	}
    return userDatas, nil
}


func (um *UserManagementService) GetRoles(ctx context.Context) ([]dataResponse.UserRolesResponse, error){
	userDao := dm.NewDao(dm.USERDAO)
	results, err := userDao.GetRoles(ctx)
	if err != nil{
		return nil, helpers.SomethingWrong
	}
	if results == nil && err == nil{
		return nil, helpers.NoresultFound
	}
	userRolesResp := []dataResponse.UserRolesResponse{}
	for results.Next(){
		userRoleResp := dataResponse.UserRolesResponse{}
		err := results.Scan(&userRoleResp.ID, &userRoleResp.RoleName, &userRoleResp.RoleDescription)
		if err != nil{
			return nil, err
		}
		userRolesResp = append(userRolesResp, userRoleResp)
	}
	return userRolesResp, nil
}

func (um *UserManagementService) CreateRoles(ctx context.Context, roles requestEntites.CreateRoles) error{
	userDao := dm.NewDao(dm.USERDAO)
	err := userDao.CreateRoles(ctx, roles)
	if err != nil{
		return err
	}
	return nil
}

func (um *UserManagementService) MigrateDBService(ctx context.Context) error{
	err := migrateDb.MigrateDb()
	if err != nil{
		return err
	}
	return nil
}
