package db

import (
	"errors"
	"pigeon/im-relation/db/model"

	"gorm.io/gorm"
)

func CreateGroup(txn *gorm.DB, group *model.GroupModel) error {
	return txn.Create(group).Error
}

func GetGroupInfo(txn *gorm.DB, groupID string) (*model.GroupModel, error) {
	var group model.GroupModel
	err := txn.Model(group).Where("group_id = ?", groupID).First(&group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &group, err
}
