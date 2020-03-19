package data

type UserDataResponse struct{
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DOB         string `json:"dob"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UserDataResponseWithRolePermission struct {
	ID int `json:"id"`
	FirstName   string `json:"first_name"`
	LastName string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Dob string `json:"dob"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
	RoleName string `json:"role_name"`
}

type UserRolesResponse struct {
	ID int `json:"id"`
	RoleName string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}



