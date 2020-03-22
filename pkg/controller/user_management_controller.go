package controller

import (
    "context"
    "fmt"
    "github.com/sirupsen/logrus"
    requestEntities "github.com/todo_list_users_service/pkg/entities/requests"
    response "github.com/todo_list_users_service/pkg/entities/responses"
    "github.com/todo_list_users_service/pkg/helpers"
    "github.com/todo_list_users_service/pkg/service"
    "net/http"
    "time"
)


type Controller struct {
    Service service.UserService
}
func (controller Controller) UserSignIn(rw http.ResponseWriter, req *http.Request) {
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("UserSignIn Startes")
    successResponse := response.Response{}
    //um := dm.NewService(dm.USERSERVICE)
    userRequest := requestEntities.UserSigninRequest{}
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    defer cancel()
    err := userRequest.PopulateUserSigninRequest(req.Body);
    if err!=nil{
        logrus.WithField("EventType", "UserSignIn").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("PopulateUserSigninRequest Failed")
        response.HandleError(rw, helpers.InvalidRequest)
    }
    err = userRequest.ValidateUserSigninRequest();
    if err!=nil{
        logrus.WithField("EventType", "UserSignIn").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("ValidateUserSigninRequest Failed")
        response.HandleError(rw, helpers.InvalidRequest)
        return
    }
    email, err := controller.Service.UserSignIn(ctx, userRequest);
    if err!= nil{
        logrus.WithField("EventType", "UserSignIn").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithError(err).Error("UserSignIn Failed")
        fmt.Println(err)
        response.HandleError(rw, err)
        return
    }
    successResponse.UserRegistration(email)
    successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "UserSignIn").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Response").Info("UserSignIn Ends")
}


func (controller Controller) UserSignUp(rw http.ResponseWriter, req *http.Request) {
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", "UserSignUp").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("UserSignUp Startes")
    successResponse := response.Response{}
    userSignUpRequest := requestEntities.UserSignupRequest{}
    //um := dm.NewService(dm.USERSERVICE)
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    defer cancel()
    err := userSignUpRequest.PopulateUserSignupRequest(req.Body);
    if err!=nil{
        logrus.WithField("EventType", "UserSignUp").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("PopulateUserSignupRequest Failed")
        response.HandleError(rw, helpers.InvalidRequest)
    }
    err = userSignUpRequest.ValidateUserSignupRequest()
    if err!=nil{
        logrus.WithField("EventType", "UserSignUp").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("ValidateUserSignupRequest Failed")
        response.HandleError(rw, helpers.InvalidRequest)
        return
    }
    token, err := controller.Service.UserSignUp(ctx, userSignUpRequest);
    if err!= nil{
        logrus.WithField("EventType", "UserSignUp").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("UserSignUp Failed")
        fmt.Println(err)
        response.HandleError(rw, err)
        return
    }
    successResponse.UserLogin(token)
    successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "UserSignUp").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("UserSignUp Ends")
}


func (controller Controller) GetUserDetails(rw http.ResponseWriter, req *http.Request){
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", "GetUserDetails").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("GetUserDetails Startes")
    successResponse := response.Response{}
    //um := dm.NewService(dm.USERSERVICE)
    emailID := req.URL.Query().Get("email")
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    defer cancel()
    getUserDetailsRequest := requestEntities.GetUserDetails{emailID}
    err := getUserDetailsRequest.ValidateUserDetails()
    if err != nil{
        logrus.WithField("EventType", "GetUserDetails").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("ValidateUserDetails Failed")
        response.HandleError(rw, err)
        return
    }
    userDatas, err := controller.Service.GetUser(ctx, getUserDetailsRequest)
    if err != nil{
        logrus.WithField("EventType", "GetUserDetails").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("GetUser Failed")
        response.HandleError(rw, err)
        return
    }
    successResponse.MakeGetUserResponse(userDatas)
    successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "GetUserDetails").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("GetUserDetails Ends")
}

func (controller Controller) CreateRoles(rw http.ResponseWriter, req *http.Request) {
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", "CreateRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("CreateRoles Startes")
    successResponse := response.Response{}
    createRole := requestEntities.CreateRoles{}
    //um := dm.NewService(dm.USERSERVICE)
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    defer cancel()
    err := createRole.PopulateCreateRoles(req.Body)
    if err != nil{
        logrus.WithField("EventType", "CreateRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("PopulateCreateRoles Failed")
        response.HandleError(rw, err)
        return
    }
    err = createRole.ValidateCreateRoles()
    if err != nil{
        logrus.WithField("EventType", "CreateRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("ValidateCreateRoles Failed")
        response.HandleError(rw, err)
        return
    }
    err = controller.Service.CreateRoles(ctx, createRole)
    if err != nil{
        logrus.WithField("EventType", "CreateRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("CreateRoles Failed")
        response.HandleError(rw, err)
        return
    }
    successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "CreateRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("CreateRoles Ends")
}

func (controller Controller) GetRoles(rw http.ResponseWriter, req *http.Request) {
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", "GetRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("GetRoles Startes")
    successResponse := response.Response{}
    //um := dm.NewService(dm.USERSERVICE)
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    defer cancel()
     resp, err := controller.Service.GetRoles(ctx)
     if err != nil{
         logrus.WithField("EventType", "GetRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
             WithError(err).Error("GetRoles Failed")
         response.HandleError(rw, err)
         return
     }
     successResponse.UserRoles(resp)
     successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "GetRoles").WithField("EventType", eventType).WithField("TraceID", traceID).
         WithField("Action","Request").Info("GetRoles Ends")
}


func (controller Controller) MigrateDB(rw http.ResponseWriter, req *http.Request) {
    eventType, traceID := getEventTypeAndTraceID(req)
    logrus.WithField("EventType", "MigrateDB").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("MigrateDB Startes")
    successResponse := response.Response{}
    //um := dm.NewService(dm.USERSERVICE)
    ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
    helpers.SetEnv()
    defer cancel()
    err := controller.Service.MigrateDBService(ctx)
    if err != nil{
        logrus.WithField("EventType", "MigrateDB").WithField("EventType", eventType).WithField("TraceID", traceID).
            WithError(err).Error("MigrateDBService Failed")
        response.HandleError(rw, err)
        return
    }
    successResponse.SendResponse(rw, http.StatusOK)
    logrus.WithField("EventType", "MigrateDB").WithField("EventType", eventType).WithField("TraceID", traceID).
        WithField("Action","Request").Info("MigrateDB Ends")
}


func getEventTypeAndTraceID(req *http.Request)(string, string){
    eventType := req.Header.Get("X-EventType")
    traceID := req.Header.Get("X-TraceID")
    return eventType, traceID
}