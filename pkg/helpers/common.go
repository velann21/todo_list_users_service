package helpers

import (
	"bytes"
	"net/http"
	"os"
	"time"
)

const (
	MYSQLCONNECTIONSTRING = "MysqlConnectionStr"
	AUTHSERVICECONNECTION = "AuthServiceConnectionStr"
)
func HttpRequest(methodType string, Url string, body []byte) (*http.Response , error) {
	req, err := http.NewRequest(methodType, Url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client http.Client
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SetEnv(){
	//os.Setenv("MYSQL_CONN", "root:root@tcp(database:3306)/UsersService?")
	os.Setenv("MYSQL_CONN", "root:Siar@123@tcp(localhost:3306)/usersrv?")
	os.Setenv("AUTHSERVICE_CONN", "http://localhost:8083/api/v1/auth/newtoken")

}

func ReadEnv(envType string) string{
	switch envType{
	case MYSQLCONNECTIONSTRING:
		return os.Getenv("MYSQL_CONN")
	case AUTHSERVICECONNECTION:
		return os.Getenv("AUTHSERVICE_CONN")
	default:
		return ""
	}
}