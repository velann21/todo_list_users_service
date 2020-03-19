package entities

type UserData struct{
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DOB         string `json:"dob"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (userData *UserData) PopulateUserData(FirstName string,LastName string,Email string,DOB string,PhoneNumber string,Password string ){
	userData.FirstName = FirstName
	userData.LastName = LastName
	userData.DOB = DOB
	userData.PhoneNumber = PhoneNumber
	userData.Email = Email
	userData.Password = Password
}
