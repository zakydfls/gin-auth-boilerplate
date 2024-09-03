package handler

import (
	"gin-auth-boilerplate/helper"
	"gin-auth-boilerplate/model/request"
	"gin-auth-boilerplate/model/response"
	"gin-auth-boilerplate/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (handler *AuthHandler) Login(ctx *gin.Context) {
	loginRequest := request.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.PanicError(err)

	accessToken, refreshToken, err_token := handler.authService.Login(loginRequest)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (handler *AuthHandler) Register(ctx *gin.Context) {
	createUser := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUser)
	helper.PanicError(err)

	handler.authService.Register(createUser)
	webResponse := response.Response{
		Code:    http.StatusCreated,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}
	ctx.JSON(http.StatusCreated, webResponse)
}
