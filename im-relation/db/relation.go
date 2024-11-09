package db

import (
	"errors"
	"pigeon/im-relation/db/model"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetAllRelationsByUsername(txn *gorm.DB, username string) ([]*model.RelationModel, error) {
	var relations []*model.RelationModel
	err := txn.Model(&model.RelationModel{}).Where("owner_id = ?", username).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func UpdateRelation(txn *gorm.DB, relation *model.RelationModel) error {
	return txn.Save(relation).Error
}

func GetRelationByUsernameAndGroupId(txn *gorm.DB, username string, groupId string) (
	*model.RelationModel, error) {
	var m model.RelationModel
	err := txn.Model(m).Where("owner_id = ? and group_id = ?", username, groupId).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &m, err
}

// 利用uniqueIndex, 先尝试插入, 如果插入失败则select for update
func InsertOrSelectForUpdateRelationByUsernameGroupId(txn *gorm.DB,
	initRow *model.RelationModel) (*model.RelationModel, error) {
	err := txn.Model(&model.RelationModel{}).Create(&initRow).Error
	if err == nil {
		return initRow, nil

	}
	if mysqlErr, ok := err.(*mysql.MySQLError); !ok || mysqlErr.Number != 1062 {
		return nil, err
	}

	var out model.RelationModel
	err = txn.Model(&model.RelationModel{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("owner_id = ? and group_id = ?",
		initRow.OwnerId, initRow.GroupId).First(&out).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func InsertRelation(txn *gorm.DB, m *model.RelationModel) error {
	return txn.Create(m).Error
}

func GetRelationByUsernameGroupId(txn *gorm.DB, username string, groupId string) (*model.RelationModel, error) {
	var m model.RelationModel
	err := txn.Model(&model.RelationModel{}).Where("owner_id = ? and group_id = ?", username, groupId).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

func SelectForUpdateRelationByUsernameGroupId(txn *gorm.DB, username string, groupId int64) (*model.RelationModel, error) {
	var m model.RelationModel
	err := txn.Model(&model.RelationModel{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("owner_id = ? and group_id = ?",
		username, groupId).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &m, err
}
