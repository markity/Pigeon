package db

import (
	"errors"
	"pigeon/im-relation/db/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllApplicationsByUsername(txn *gorm.DB, username string) ([]*model.ApplyModel, error) {
	var applications []*model.ApplyModel
	err := txn.Model(&model.ApplyModel{}).Where("owner_id = ?", username).Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

// InsertApply 插入申请
func InsertApply(txn *gorm.DB, apply *model.ApplyModel) (inserted bool, err error) {
	err = txn.Create(apply).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return false, nil
	}

	return false, err
}

func SelectForUpdateApplyByUsername(txn *gorm.DB, username string) (*model.ApplyModel, error) {
	var apply model.ApplyModel
	err := txn.Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.ApplyModel{}).Where("owner_id = ?", username).First(&apply).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &apply, nil
}

func UpdateApply(txn *gorm.DB, apply *model.ApplyModel) error {
	return txn.Save(apply).Error
}
