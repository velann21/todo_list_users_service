package routes

import ("context"
	proto "github.com/velann21/todo_list_users_service/pkg/proto"
)

type ServerRoutes struct {
}


func (server ServerRoutes) RegisterUser(ctx context.Context, request *proto.UserRegistrationRequest)(*proto.UserRegistrationResponse, error){

	return nil, nil
}


func (server ServerRoutes) LoginUser(context.Context, *proto.UserRegistrationRequest) (*proto.UserRegistrationResponse, error){

	return nil, nil
}


func (server ServerRoutes) GetUserDetails(context.Context, *proto.UserRegistrationRequest) (*proto.UserRegistrationResponse, error){

	return nil, nil
}
