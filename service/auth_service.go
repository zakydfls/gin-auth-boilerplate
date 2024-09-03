package service

import "gin-auth-boilerplate/model/request"

type AuthService interface {
	Login(user request.LoginUserRequest) (string, string, error)
	Register(user request.CreateUserRequest)
}
