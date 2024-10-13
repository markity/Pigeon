package db

import (
	"errors"
	"pigeon/im-relation/db/model"

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

// 利用uniqueIndex, 先尝试插入, 如果插入失败则select for update
func InsertOrSelectForUpdateRelationByUsernameGroupId(txn *gorm.DB,
	initRow *model.RelationModel) (*model.RelationModel, error) {
	err := txn.Model(&model.RelationModel{}).Create(&initRow).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return initRow, nil
	}

	var out model.RelationModel
	err = txn.Model(&model.RelationModel{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("owner_id = ? and group_id = ?",
		initRow.OwnerId, initRow.GroupId).First(&out).Error
	if err != nil {
		return nil, err
	}
	return &out, nil
}
