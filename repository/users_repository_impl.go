package repository

import (
	"errors"
	"gin-auth-boilerplate/helper"
	"gin-auth-boilerplate/model/entity"
	"gin-auth-boilerplate/model/request"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (u *UserRepositoryImpl) Create(users entity.Users) {
	result := u.Db.Create(&users)
	helper.PanicError(result.Error)
}
func (u *UserRepositoryImpl) Update(users entity.Users) {
	var updateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.Db.Model(&users).Updates(updateUsers)
	helper.PanicError(result.Error)
}
func (u *UserRepositoryImpl) Delete(userId int) {
	var users entity.Users
	result := u.Db.Where("id = ?", userId).Delete(&users)
	helper.PanicError(result.Error)
}
func (u *UserRepositoryImpl) FindById(userId int) (entity.Users, error) {
	var users entity.Users
	result := u.Db.Find(&users, userId)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("user is not found")
	}

}
func (u *UserRepositoryImpl) FindByUsername(username string) (entity.Users, error) {
	var users entity.Users
	result := u.Db.Find(&users, "username = ?", username)
	if result != nil {
		return users, nil
	} else {
		return users, errors.New("invalid username")
	}

}
func (u *UserRepositoryImpl) FindAll() []entity.Users {
	var users []entity.Users
	result := u.Db.Find(&users)
	helper.PanicError(result.Error)
	return users
}
