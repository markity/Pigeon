package db

import (
	"crypto/sha256"
	"errors"
	"pigeon/im-auth-route/db/model"

	"gorm.io/gorm"
)

func ToSha256(bs []byte) []byte {
	s := sha256.New()
	s.Write(bs)
	s.Sum(nil)
	return s.Sum(nil)
}

func (db *DB) GetUserByUsername(username string) (*model.UserModel, error) {
	m := &model.UserModel{}
	err := db.db.Model(&model.UserModel{}).Where("username = ?", username).First(m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
