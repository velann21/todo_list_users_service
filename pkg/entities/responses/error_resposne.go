package entities

import (
"encoding/json"
"net/http"

helpers "github.com/velann21/todo_list_users_service/pkg/helpers"
)

// ErrorResponse response for error
type ErrorResponse struct {
Data    []interface{} `json:"data"`
Status  string        `json:"status"`
Message string        `json:"message"`
}

// HandleError handles error and send response
func HandleError(rw http.ResponseWriter, err error) {
// build default response
var response *ErrorResponse
response = &ErrorResponse{Data: make([]interface{}, 0), Message: "somethingWentWrong",
Status: http.StatusText(http.StatusInternalServerError)}
rw.Header().Set("Content-Type", "application/json")
// set header, message and status
switch err {
case helpers.ErrUserNotFound:
rw.WriteHeader(http.StatusBadRequest)
response.Message = "invalidUserId"
response.Status = http.StatusText(http.StatusBadRequest)
case helpers.UserAlreadyExist:
	rw.WriteHeader(http.StatusBadRequest)
	response.Message = "UserAlreadyExists"
	response.Status = http.StatusText(http.StatusBadRequest)
case helpers.InvalidRequest:
rw.WriteHeader(http.StatusBadRequest)
response.Message = "invalidRequest"
response.Status = http.StatusText(http.StatusBadRequest)
default:
rw.WriteHeader(http.StatusInternalServerError)
}

// send response
json.NewEncoder(rw).Encode(response)
return
}
