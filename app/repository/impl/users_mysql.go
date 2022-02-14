package impl

import (
	"gorm.io/gorm"
	"todolist-assignment/app/domain"
	"todolist-assignment/app/repository"
	"todolist-assignment/infra/errors"
)

type users struct {
	DB *gorm.DB
}

func NewMysqlUsersRepository(db *gorm.DB) repository.IUsers {
	return &users{
		DB: db,
	}
}

func (uRep *users) Save(user *domain.User) (*domain.User, *errors.RestErr) {

	res := uRep.DB.Model(&domain.User{}).Create(&user)

	if res.Error != nil {
		return nil, errors.NewBadRequestError("username or email already used")
	}
	return user, nil
}

func (uRep *users) AuthAndGetID(username string, password string) (uint, *errors.RestErr) {
	user := domain.User{}
	res := uRep.DB.Model(&domain.User{}).Where("username = ? and password = ?", username, password).Find(&user)
	if res.Error != nil {
		return 0, errors.NewUnauthorizedError(res.Error.Error())
	}
	if res.RowsAffected > 0 {
		return user.ID, nil
	}
	return 0, errors.NewUnauthorizedError("not found")
}
