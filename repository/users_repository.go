package repository

import "gin-auth-boilerplate/model/entity"

type UsersRepository interface {
	Create(users entity.Users)
	Update(users entity.Users)
	Delete(userId int)
	FindById(userId int) (entity.Users, error)
	FindAll() []entity.Users
	FindByUsername(username string) (entity.Users, error)
}
