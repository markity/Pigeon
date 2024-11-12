package db

import (
	"errors"
	"slices"

	"pigeon/im-relay/db/model"

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

func GetMessages(txn *gorm.DB, groupId string, maxSeqId int64, limit int64) ([]*model.MessageModel, error) {
	var result []*model.MessageModel
	err := txn.Model(model.MessageModel{}).Where("seq_id <= ?", maxSeqId).Order("seq_id desc").Limit(int(limit)).Find(&result).Error
	slices.Reverse(result)
	return result, err
}
