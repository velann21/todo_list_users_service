package ipc_responses

import (
	"encoding/json"
	"github.com/todo_list_users_service/pkg/helpers"
	"io"
)

type NewTokenResponse struct {
	Status string `json:"status"`
	Data []Data `json:"data"`

}

type Data struct {
	Token string`json:"Token"`
}

func (newTokenResp *NewTokenResponse) PopulateNewTokenResp(body io.ReadCloser) (*NewTokenResponse, error){
	decoder := json.NewDecoder(body)
	err := decoder.Decode(newTokenResp)
	if err != nil {
		return nil, helpers.SomethingWrong
	}
	return newTokenResp, nil
}
