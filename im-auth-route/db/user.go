package db

import (
	"pigeon/im-auth-route/db/model"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, user *model.UserModel) error {
	return db.Create(user).Error
}
