package db

import (
	"pigeon/im-chat-evloop/db/model"

	"gorm.io/gorm"
)

func InsertMessage(txn *gorm.DB, m *model.MessageModel) error {
	return txn.Create(m).Error
}
