package entities

import (
	"encoding/json"
	"github.com/todo_list_users_service/pkg/helpers"
	"io"
	"log"
)

type UserSigninRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DOB         string `json:"dob"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Role        []int    `json:"role"`
	Password    string `json:"password"`
}

func (userSignin *UserSigninRequest) PopulateUserSigninRequest(body io.ReadCloser) error{
	decoder := json.NewDecoder(body)
	err := decoder.Decode(userSignin)
	if err != nil {
		return helpers.NotValidRequestBody
	}
	return nil
}

func (userSignin *UserSigninRequest) ValidateUserSigninRequest() error{
	if userSignin.Email == "" || userSignin.Email == " "{
		return helpers.NotValidRequestBody
	}

	if userSignin.FirstName == "" || userSignin.FirstName == " "{
		return helpers.NotValidRequestBody
	}

	if len(userSignin.Role) <= 0{
		return helpers.NotValidRequestBody
	}

	for _, value :=range userSignin.Role{
        if !helpers.Roles_checker(value){
			return helpers.NotValidRequestBody
		}
	}

	if len(userSignin.Password) <= 5 || userSignin.Password == "" || userSignin.Password == " "{
		return helpers.InvalidPassword
	}

	if userSignin.PhoneNumber == "" || userSignin.PhoneNumber == " "{
		return helpers.NotValidRequestBody
	}
	return nil
}

func (userSignin *UserSigninRequest) MarshalUserSigninRequest() ([]byte, error) {
	var jsonData []byte
	jsonData, err := json.Marshal(userSignin)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

type UserSignupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (userSignup *UserSignupRequest) PopulateUserSignupRequest(body io.ReadCloser) error{
	decoder := json.NewDecoder(body)
	err := decoder.Decode(userSignup)
	if err != nil {
		return helpers.NotValidRequestBody
	}
	return nil
}

func (userSignup *UserSignupRequest) ValidateUserSignupRequest() error{
	if userSignup.Email == "" || userSignup.Email == " "{
		return helpers.NotValidRequestBody
	}

	if len(userSignup.Password) <= 5 || userSignup.Password == "" || userSignup.Password == " "{
		return helpers.InvalidPassword
	}
	return nil
}

func (userSignup *UserSignupRequest) MarshalUserSignupRequest() ([]byte, error) {
	var jsonData []byte
	jsonData, err := json.Marshal(userSignup)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return jsonData, nil
}

type GetUserDetails struct {
     EmailID string
}

func (getUserDetails *GetUserDetails) ValidateUserDetails() error{
	if getUserDetails.EmailID == "" || getUserDetails.EmailID == " "{
		return helpers.InvalidRequest
	}
	return nil
}


type CreateRoles struct {
	RoleName string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

func (createRoles *CreateRoles) PopulateCreateRoles(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(createRoles)
	if err != nil{
		return helpers.NotValidRequestBody
	}
    return nil
}

func (createRoles *CreateRoles) ValidateCreateRoles() error {
	if createRoles.RoleName == "" || createRoles.RoleName == " "{
		return helpers.NotValidRequestBody
	}

	if createRoles.RoleDescription == "" || createRoles.RoleDescription == " "{
		return helpers.NotValidRequestBody
	}
	return nil
}


