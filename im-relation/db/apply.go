package db

import (
	"errors"
	"pigeon/im-relation/db/model"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllApplicationsByUsername(txn *gorm.DB, username string) ([]*model.ApplyModel, error) {
	var applications []*model.ApplyModel
	err := txn.Model(&model.ApplyModel{}).Where("group_owner_id = ?", username).Find(&applications).Error
	if err != nil {
		return nil, err
	}

	return applications, nil
}

func GetApplyByUsernameAndGroupId(txn *gorm.DB, username string, groupId string) (*model.ApplyModel, error) {
	var m model.ApplyModel
	err := txn.Model(m).Where("owner_id = ? and group_id = ?", username, groupId).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

// InsertApply 插入申请
func InsertApply(txn *gorm.DB, apply *model.ApplyModel) (inserted bool, err error) {
	err = txn.Create(apply).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return false, nil
	}

	return false, err
}

func UpdateApply(txn *gorm.DB, apply *model.ApplyModel) error {
	return txn.Save(apply).Error
}

// 利用uniqueIndex, 先尝试插入, 如果插入失败则select for update
func InsertOrSelectForUpdateApplyByUsernameGroupId(txn *gorm.DB,
	initRow *model.ApplyModel) (*model.ApplyModel, error) {
	err := txn.Model(&model.ApplyModel{}).Create(&initRow).Error
	if err == nil {
		return initRow, nil

	}
	if mysqlErr, ok := err.(*mysql.MySQLError); !ok || mysqlErr.Number != 1062 {
		return nil, err
	}

	var out model.ApplyModel
	err = txn.Model(&model.ApplyModel{}).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("owner_id = ? and group_id = ?", initRow.OwnerId, initRow.GroupId).
		First(&out).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}
