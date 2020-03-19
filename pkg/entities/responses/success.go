package entities

import (
	"encoding/json"
	"github.com/todo_list_users_service/pkg/entities/data"
	"net/http"

)

// Response struct
type Response struct {
	Status string                   `json:"status"`
	Data   []map[string]interface{} `json:"data"`
	Meta   map[string]interface{}   `json:"meta,omitempty"`
}

func (entity *Response)  UserRegistration(name *string){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["UserName"] = name
	responseData = append(responseData, data)
    entity.Data = responseData
    metaData := make(map[string]interface{})
    metaData["message"] = "User created"
}

func (entity *Response)  UserLogin(token *string){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["token"] = token
	responseData = append(responseData, data)
	entity.Data = responseData
}

func (entity *Response) MakeGetUserResponse(userData []data.UserDataResponseWithRolePermission){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	roles := make([]int,0)

	for _,V := range userData{
		data["firstName"] = V.FirstName
		data["lastName"] = V.LastName
		data["email"] = V.Email
		data["phoneNumber"] = V.PhoneNumber
		data["dob"] = V.Dob
		roles = append(roles, V.RoleID)
	}
	data["roles"] = roles
	responseData = append(responseData, data)
	entity.Data = responseData
}

func (entity *Response)  UserRoles(roles []data.UserRolesResponse){
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["roles"] = roles
	responseData = append(responseData, data)
	entity.Data = responseData
}

// SendResponse send http response
func (entity *Response) SendResponse(rw http.ResponseWriter, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		entity.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		entity.Status = http.StatusText(http.StatusAccepted)
	default:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	}

	// send response
	json.NewEncoder(rw).Encode(entity)
	return
}