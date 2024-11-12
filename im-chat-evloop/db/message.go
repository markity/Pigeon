package db

import (
	"errors"
	"pigeon/im-chat-evloop/db/model"

	"gorm.io/gorm"
)

func InsertMessage(txn *gorm.DB, m *model.MessageModel) error {
	return txn.Create(m).Error
}

func GetMessageByIdempotentKey(txn *gorm.DB, key string) (*model.MessageModel, error) {
	var m model.MessageModel
	err := txn.Model(model.MessageModel{}).Where("idempotent_key = ?", key).First(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
