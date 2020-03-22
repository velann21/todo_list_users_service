package ipc_responses

import (
	"encoding/json"
	"log"
)

type AuthRequestBody struct {
	FirstName   string `json:"userName"`
	Email       string `json:"emailID"`
	RoleID []int `json:"roles"`
	PermissionID []int `json:"permissions"`
	PermissionName []string `json:"permissionsName"`
}

func (authRequestBody *AuthRequestBody) MarshalAuthRequestBody()([]byte, error){
	var jsonData []byte
	jsonData, err := json.Marshal(authRequestBody)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}


