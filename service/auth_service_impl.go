package service

import (
	"errors"
	"gin-auth-boilerplate/config"
	"gin-auth-boilerplate/helper"
	"gin-auth-boilerplate/model/entity"
	"gin-auth-boilerplate/model/request"
	"gin-auth-boilerplate/repository"
	"gin-auth-boilerplate/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewAuthServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(user request.LoginUserRequest) (string, string, error) {
	new_user, user_err := a.UsersRepository.FindByUsername(user.Username)
	if user_err != nil {
		return "", "", errors.New("creds not found")
	}
	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, user.Password)
	if verify_error != nil {
		return "", "", errors.New("creds not found")
	}
	// generate access token
	accessClaims := entity.CustomClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	accessToken, token_err := utils.GenerateToken(accessClaims, config.TokenSecret)
	if token_err != nil {
		return "", "", errors.New("error generating session")
	}
	// generate refresh token
	refreshToken, refresh_err := utils.GenerateRefreshToken(user.Username, config.RefreshSecret)
	if refresh_err != nil {
		return "", "", errors.New("error generating session")
	}
	return accessToken, refreshToken, nil
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(user request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(user.Password)
	helper.PanicError(err)

	newUser := entity.Users{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	a.UsersRepository.Create(newUser)
}
