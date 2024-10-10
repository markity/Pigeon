package db

import (
	"errors"
	"pigeon/im-relation/db/model"

	"gorm.io/gorm"
)

func CreateGroup(txn *gorm.DB, group *model.GroupModel) error {
	return txn.Create(group).Error
}

func GetGroupInfo(txn *gorm.DB, groupID int64) (*model.GroupModel, error) {
	var group model.GroupModel
	err := txn.Where("id = ?", group).First(&group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &group, err
}
