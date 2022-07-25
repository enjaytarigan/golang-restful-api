package repository

import "brodo-demo/entity"

type UserRepository interface {
	Insert(user entity.User) (userId int, err error)
	VerifyAvailableUsername(username string) bool
	FindByUsername(username string) (entity.User, error)
	VerifyUserIsExist(userId int) error
}
